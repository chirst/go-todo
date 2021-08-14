package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/chirst/go-todo/auth"
	"github.com/chirst/go-todo/config"
	"github.com/chirst/go-todo/database"
	"github.com/chirst/go-todo/http/rest"
	"github.com/chirst/go-todo/todo"
	"github.com/chirst/go-todo/user"

	redoc "github.com/go-openapi/runtime/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

func main() {
	inMemoryFlag := flag.Bool("use-memory", false, "use a temporary database")
	flag.Parse()

	var todosRepo todo.Repository
	var usersRepo user.Repository
	if *inMemoryFlag {
		log.Println("using in memory db")
		todosRepo = new(todo.MemoryRepository)
		usersRepo = new(user.MemoryRepository)
	} else {
		db := database.InitDB()
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
		r.Patch("/todos/{todoID}/complete", rest.CompleteTodo(todoService))
		r.Delete("/todos/{todoID}", rest.DeleteTodo(todoService))
	})

	// unprotected routes
	router.Group(func(r chi.Router) {
		r.Post("/users", rest.AddUser(usersService))
		r.Post("/login", rest.Login(usersService))

		// docs
		sh := redoc.Redoc(redoc.RedocOpts{
			SpecURL: "/swagger.yaml",
			Title:   "Todo API Documentation",
		}, nil)
		r.Handle("/docs", sh)
		r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	})

	address := config.ServerAddress()
	log.Printf("server listening on %s\n", address)
	http.ListenAndServe(address, router)
}
