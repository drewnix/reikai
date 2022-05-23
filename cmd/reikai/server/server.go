package server

import (
	"fmt"
	"net/http"

	"github.com/heptiolabs/healthcheck"
	"github.com/urfave/negroni"
)

func setupRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to reikai!")
	})
	// Healthcheck
	health := healthcheck.NewHandler()
	mux.Handle("/live", health)
	mux.Handle("/ready", health)

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)

	return n
}

func Serve() error {
	n := setupRoutes()

	http.ListenAndServe(":3000", n)

	return nil
}
