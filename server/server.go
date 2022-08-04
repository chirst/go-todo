package server

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/chirst/go-todo/auth"
	"github.com/chirst/go-todo/http/rest"
	"github.com/chirst/go-todo/todo"
	"github.com/chirst/go-todo/user"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	redoc "github.com/go-openapi/runtime/middleware"
)

// GetRouter configures the application router
func GetRouter(db *sql.DB) http.Handler {
	var todosRepo todo.Repository
	var usersRepo user.Repository
	if db != nil {
		todosRepo = &todo.PostgresRepository{DB: db}
		usersRepo = &user.PostgresRepository{DB: db}
	} else {
		log.Println("using in memory db")
		todosRepo = new(todo.MemoryRepository)
		usersRepo = new(user.MemoryRepository)
	}
	todoService := todo.NewService(todosRepo)
	usersService := user.NewService(usersRepo)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(httprate.LimitByIP(100, time.Minute))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Protected routes
	router.Group(func(r chi.Router) {
		r.Use(auth.Verifier)
		r.Use(auth.Authenticator)

		r.Get("/todos", rest.GetTodos(todoService))
		r.Post("/todos", rest.AddTodo(todoService))
		r.Patch("/todos/{todoID}", rest.PatchTodo(todoService))
		r.Delete("/todos/{todoID}", rest.DeleteTodo(todoService))
		r.Get("/priorities", rest.GetPriorities(todoService))
	})

	// Unprotected routes
	router.Group(func(r chi.Router) {
		r.Post("/users", rest.AddUser(usersService))
		r.Post("/login", rest.Login(usersService))

		// Docs
		sh := redoc.Redoc(redoc.RedocOpts{
			SpecURL: "/swagger.yaml",
			Title:   "Todo API Documentation",
		}, nil)
		r.Handle("/docs", sh)
		r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	})

	return (http.Handler)(router)
}
