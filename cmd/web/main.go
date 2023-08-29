package main

import (
	"chatapp/internal/handlers"
	"log"
	"net/http"
)

func main() {

	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Starting server on :4000")

	_ = http.ListenAndServe(":4000", mux)
}
