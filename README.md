# flexhttp

<a href="https://pkg.go.dev/github.com/go-flexible/flexhttp"><img src="https://pkg.go.dev/badge/github.com/go-flexible/flexhttp.svg" alt="Go Reference"></a>

A [flex](https://github.com/go-flexible/flex) compatible http server.

```go
// Create some router and endpoints...
router := http.NewServeMux()
router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
        fmt.Fprint(rw, "Hello, world!\n")
})

// Create a standard http server.
srv := &http.Server{
        Addr:              ":8080",
        Handler:           router,
        ReadTimeout:       5 * time.Second, 
        ReadHeaderTimeout: time.Second,
        // Missing timeouts will be set to a sane default.
}

// Run it, or better yet, let `flex` run it!
flexhttp.New(srv).Run(context.Background())
```
