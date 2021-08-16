package flexhttp_test

import (
	"context"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/go-flexible/flexhttp"
)

var (
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	ctx context.Context
	now time.Time
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	now = time.Now()
	os.Exit(m.Run())
}

func Example() {
	server := flexhttp.New(&http.Server{
		Handler:     handler,
		ReadTimeout: 1 * time.Second,
	})

	go func() {
		_ = server.Run(ctx)
	}()
}

func TestNewHTTPServer(t *testing.T) {
	testcases := []struct {
		name            string
		srv             *http.Server
		expectedTimeout time.Duration
	}{
		{
			name: "nil server must not fail",
			srv:  nil,
		},
		{
			name: "defaults are used when no timeouts are provided",
			srv:  &http.Server{},
		},
		{
			name: "defaults are overridden when user chooses a timeout",
			srv: &http.Server{
				ReadTimeout:       time.Minute,
				ReadHeaderTimeout: time.Minute,
				WriteTimeout:      time.Minute,
				IdleTimeout:       time.Minute,
			},
			expectedTimeout: time.Minute,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			server := flexhttp.New(tt.srv)
			if server.Server == nil {
				t.Error("http server must not be nil")
			}
			if tt.expectedTimeout > 0 {
				equal(t, server.ReadTimeout, tt.expectedTimeout)
				equal(t, server.ReadHeaderTimeout, tt.expectedTimeout)
				equal(t, server.WriteTimeout, tt.expectedTimeout)
				equal(t, server.IdleTimeout, tt.expectedTimeout)
			}
			notEqual(t, server.ReadTimeout, 0)
			notEqual(t, server.ReadHeaderTimeout, 0)
			notEqual(t, server.WriteTimeout, 0)
			notEqual(t, server.IdleTimeout, 0)
		})
	}
}

func equal(t *testing.T, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got: %#[1]v (%[1]T), but wanted: %#[2]v (%[2]T)", got, want)
	}
}
func notEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		t.Fatalf("got: %#[1]v (%[1]T), but wanted: %#[2]v (%[2]T)", got, want)
	}
}
