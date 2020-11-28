package main

import (
	"fmt"
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
	fmt.Println("started server on localhost:3000")
	http.ListenAndServe("localhost:3000", router)
}
