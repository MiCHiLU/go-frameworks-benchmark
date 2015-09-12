package main

import (
	"fmt"
	"net/http"
)

func init() {
	calcMem("http", initHttp)
}

func initHttp() {
	h := http.NewServeMux()
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, World")
	})
	registerHandler("http", h)
}
