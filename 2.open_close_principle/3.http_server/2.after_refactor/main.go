package main

import (
	"net/http"
)

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func healthCheck(resp http.ResponseWriter, _ *http.Request) {
	resp.WriteHeader(http.StatusNoContent)
}

type CustomHandlerFunc func(w http.ResponseWriter, r *http.Request)

func (f CustomHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func healthCheckUsage() {
	http.Handle("/health", CustomHandlerFunc(healthCheck))
}

func main() {
	healthCheckUsage()
}
