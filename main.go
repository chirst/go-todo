package main

import (
	"database/sql"
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

	// TODO: extract to config and make repo configurable
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=12345 dbname=todo sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// todosRepo := new(todo.MemoryRepository)
	todosRepo := &todo.PostgresRepository{DB: db}
	// usersRepo := new(user.MemoryRepository)
	usersRepo := &user.PostgresRepository{DB: db}
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
	fmt.Printf("starting server on %s\n", address)
	http.ListenAndServe(address, router)
}
