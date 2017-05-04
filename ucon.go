package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/favclip/ucon"
	"golang.org/x/net/context"
)

func init() {
	calcMem("ucon", initUcon)
}

func initUcon() {
	ucon.Middleware(ucon.HTTPRWDI())
	ucon.Middleware(ucon.NetContextDI())
	ucon.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request, _ context.Context) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		io.WriteString(w, "Hello, World")
	})
	ucon.HandleFunc("GET", "/{name}", func(w http.ResponseWriter, r *http.Request, ctx context.Context) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello, %s", ctx.Value(ucon.PathParameterKey).(map[string]string)["name"])
	})
	ucon.DefaultMux.Prepare()
	registerHandler("ucon", ucon.DefaultMux)
}
