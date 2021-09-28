package cmd

import (
	"context"
	"github.com/ksbomj/sfkp/services/order/bus"
	"os"
	"os/signal"
	"syscall"

	"github.com/ksbomj/sfkp/services/order/api"
)

// ServerCommand with command line flags and env
type ServerCommand struct {
	Port    int    `long:"port" env:"PORT" default:"80" description:"port"`
	Address string `long:"address" env:"ADDRESS" default:"" description:"listening address"`

	BrokerAddress string `long:"broker" env:"BROKER_ADDRESS" default:"broker:9092" description:"Message broker address"`

	CommonOpts
}

// server wraps everything needed for server
type server struct {
	*ServerCommand
	restSrv *api.RestApi

	terminated chan struct{}
}

// Execute is the entry point for "server" command, called by flag parser
func (s *ServerCommand) Execute(_ []string) error {
	s.Logger.Printf("starting server on %s:%d", s.Address, s.Port)

	ctx, cancel := context.WithCancel(context.Background())

	// catch signal and invoke graceful termination
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		s.Logger.Printf("interrupt signal")
		cancel()
	}()

	app, err := s.newServer(ctx)

	if err != nil {
		s.Logger.Printf("failed to create a server, %+v", err)
		return err
	}

	if err = app.run(ctx); err != nil {
		s.Logger.Printf("terminated with error %+v", err)
		return err
	}

	s.Logger.Printf("service terminated")

	return nil
}

// newServer builds a server
func (s *ServerCommand) newServer(ctx context.Context) (*server, error) {

	messageBus := bus.NewBus(s.BrokerAddress)

	api := &api.RestApi{
		Logger: s.Logger,
		MessageBus: messageBus,
	}

	return &server{
		ServerCommand: s,
		restSrv:       api,

		terminated: make(chan struct{}),
	}, nil
}

func (srv *server) run(ctx context.Context) error {
	go func() {
		// shutdown on context cancellation
		<-ctx.Done()
		srv.ServerCommand.Logger.Print("shutdown initiated")
		srv.restSrv.Shutdown()
	}()

	srv.restSrv.Run(srv.Address, srv.Port)

	close(srv.terminated)
	return nil
}

// Wait application completion
func (srv *server) Wait() {
	<-srv.terminated
}
