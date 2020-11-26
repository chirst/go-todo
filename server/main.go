package main

import (
	"net/http"
	"todo/adding"
	"todo/http/rest"
	"todo/listing"
	"todo/persistence/memory"
)

func main() {
	todosRepo := new(memory.Storage)
	listingService := listing.NewService(todosRepo)
	addingService := adding.NewService(todosRepo)
	router := rest.Handler(listingService, addingService)
	http.ListenAndServe(":3000", router)
}
