package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-flexible/flex"
	"github.com/go-flexible/flexhttp"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "hello, world\n")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	flex.MustStart(
		context.Background(),
		flexhttp.NewHTTPServer(srv),
	)
}
