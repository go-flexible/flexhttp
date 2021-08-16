package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-flexible/flexhttp"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello, world!\n")
	})

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadTimeout:       time.Second,
		ReadHeaderTimeout: time.Second,
		WriteTimeout:      time.Second,
		IdleTimeout:       time.Second,
	}

	flexhttp.New(srv).Run(context.Background())
}
