package main

import (
	"net/http"
	"todo/http/rest"
	"todo/listing"
	"todo/persistence/memory"
)

func main() {
	todosRepo := new(memory.Storage)
	listingService := listing.NewService(todosRepo)
	router := rest.Handler(listingService)
	http.ListenAndServe(":3000", router)
}
