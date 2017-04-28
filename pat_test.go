package main

import (
	"net/http"
	"testing"
)

func TestPatParam(t *testing.T) {
	testRequestWithPathParam(t, getHandler("pat"))
}

func BenchmarkPatSimple(b *testing.B) {
	req, _ := http.NewRequest("GET", "/", nil)
	benchRequest(b, getHandler("pat"), req)
}

func BenchmarkPatParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/gopher", nil)
	benchRequest(b, getHandler("pat"), req)
}
