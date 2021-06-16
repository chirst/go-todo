package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"time"
	"todo/auth"
	"todo/config"
	"todo/http/rest"
	"todo/todo"
	"todo/user"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

func main() {
	config.InitConfig()
	inMemoryFlag := flag.Bool("use-memory", false, "enable to swap database for in memory")
	flag.Parse()

	var todosRepo todo.Repository
	var usersRepo user.Repository
	if *inMemoryFlag {
		todosRepo = new(todo.MemoryRepository)
		usersRepo = new(user.MemoryRepository)
	} else {
		// TODO: migrations
		db, err := sql.Open("postgres", config.GetPostgresSourceName())
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		todosRepo = &todo.PostgresRepository{DB: db}
		usersRepo = &user.PostgresRepository{DB: db}
	}
	todoService := todo.NewService(todosRepo)
	usersService := user.NewService(usersRepo)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(httprate.LimitByIP(100, time.Minute))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// protected routes
	router.Group(func(r chi.Router) {
		r.Use(auth.Verifier)
		r.Use(auth.Authenticator)

		r.Get("/todos", rest.GetTodos(todoService))
		r.Post("/todos", rest.AddTodo(todoService))
	})

	// unprotected routes
	router.Group(func(r chi.Router) {
		r.Post("/users", rest.AddUser(usersService))
		r.Post("/login", rest.Login(usersService))
	})

	address := config.GetAddress()
	fmt.Printf("server listening on %s\n", address)
	http.ListenAndServe(address, router)
}
