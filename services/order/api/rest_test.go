package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func startupT(t *testing.T) (ts *httptest.Server, rpi *RestApi, teardown func()) {
	srv := &RestApi{}

	ts = httptest.NewServer(srv.router())

	teardown = func() {
		ts.Close()
	}

	return ts, srv, teardown
}

func get(t *testing.T, url string) (response string, statusCode int) {
	r, err := http.Get(url)
	require.NoError(t, err)
	body, err := ioutil.ReadAll(r.Body)
	require.NoError(t, err)
	require.NoError(t, r.Body.Close())
	return string(body), r.StatusCode
}

func TestRestApi_Health(t *testing.T) {
	ts, _, teardown := startupT(t)
	defer teardown()

	_, code := get(t, ts.URL+"/health")
	assert.Equal(t, 200, code)
}

func TestRestApi_Shutdown(t *testing.T) {
	logger := log.New(os.Stdout, "TEST", 0)

	srv := RestApi{
		Logger: logger,
	}
	done := make(chan bool)

	// without waiting for channel close at the end goroutine will stay alive after test finish
	// which would create data race with next test
	go func() {
		time.Sleep(200 * time.Millisecond)
		srv.Shutdown()
		close(done)
	}()

	st := time.Now()
	srv.Run("127.0.0.1", 0)
	assert.True(t, time.Since(st).Seconds() < 1, "should take about 100ms")
	<-done
}
