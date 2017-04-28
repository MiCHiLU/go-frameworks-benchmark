package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bmizerany/pat"
)

func init() {
	calcMem("pat", initPat)
}

func initPat() {
	h := pat.New()
	h.Get("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		io.WriteString(w, "Hello, World")
	}))
	h.Get("/:name", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, %s", r.URL.Query().Get(":name"))
	}))
	registerHandler("pat", h)
}
