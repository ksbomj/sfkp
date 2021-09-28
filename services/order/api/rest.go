package api

import (
	"context"
	"fmt"
	"github.com/ksbomj/sfkp/services/order/bus"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
)

type RestApi struct {
	Version string
	Logger  *log.Logger
	MessageBus *bus.Bus

	httpServer *http.Server
	lock       sync.Mutex
}

func (rpi *RestApi) Run(address string, port int) {
	if address == "*" {
		address = ""
	}

	rpi.Logger.Printf("running http the rest server on %s:%d", address, port)

	rpi.lock.Lock()
	rpi.httpServer = rpi.makeHTTPServer(address, port, rpi.router())
	rpi.lock.Unlock()

	err := rpi.httpServer.ListenAndServe()
	rpi.Logger.Printf("http server terminated, %s", err)
}

// Shutdown rest http server
func (rpi *RestApi) Shutdown() {
	rpi.Logger.Print("shutdown rest server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rpi.lock.Lock()
	if rpi.httpServer != nil {
		if err := rpi.httpServer.Shutdown(ctx); err != nil {
			rpi.Logger.Printf("http shutdown error, %s", err)
		}
		rpi.Logger.Print("shutdown http server completed")
	}

	rpi.lock.Unlock()
}

func (rpi *RestApi) makeHTTPServer(address string, port int, router http.Handler) *http.Server {
	return &http.Server{
		Addr:              fmt.Sprintf("%s:%d", address, port),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       15 * time.Second,
	}
}

func (rpi *RestApi) router() chi.Router {
	r := chi.NewRouter()
	r.Get("/health", rpi.health)

	r.Post("/orders", rpi.storeOrder)

	return r
}

func (rpi *RestApi) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (rpi *RestApi) storeOrder(w http.ResponseWriter, r *http.Request) {

}
