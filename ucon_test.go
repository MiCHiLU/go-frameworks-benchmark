package main

import (
	"net/http"
	"testing"
)

func TestUconParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("ucon"))
}

func BenchmarkUconSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("ucon"), req)
}

func BenchmarkUconParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher", nil)
	benchRequest(b, getHandler("ucon"), req)
}
