package main

import "net/http"

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type CustomHandler struct{}

func (c *CustomHandler) ServeHTTP(resp http.ResponseWriter, _ *http.Request) {
	resp.WriteHeader(http.StatusNoContent)
}

var handler *CustomHandler

func healthCheckUage() {
	http.Handle("/health", handler)
}

func main() {
	healthCheckUage()
}
