package main

import (
	"chatapp/internal/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

// routes returns a http.Handler with all application routes defined.
func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
