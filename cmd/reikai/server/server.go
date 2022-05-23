package server

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func Serve() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to reikai!")
	})

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)

	return nil
}
