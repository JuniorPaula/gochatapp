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

	return mux
}
