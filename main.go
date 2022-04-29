package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/chirst/go-todo/config"
	"github.com/chirst/go-todo/database"
	"github.com/chirst/go-todo/server"
)

func main() {
	inMemoryFlag := flag.Bool("use-memory", false, "use a temporary database")
	flag.Parse()

	var db *sql.DB
	if !*inMemoryFlag {
		db = database.InitDB()
		defer db.Close()
	}
	router := server.GetRouter(db)

	address := config.ServerAddress()
	log.Printf("server listening on %s\n", address)
	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Fatal(err)
	}
}
