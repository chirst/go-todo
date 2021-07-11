package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
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
	inMemoryFlag := flag.Bool("use-memory", false, "use a temporary database")
	flag.Parse()

	var todosRepo todo.Repository
	var usersRepo user.Repository
	if *inMemoryFlag {
		fmt.Println("using in memory db")
		todosRepo = new(todo.MemoryRepository)
		usersRepo = new(user.MemoryRepository)
	} else {
		db := initDB()
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

	address := config.ServerAddress()
	fmt.Printf("server listening on %s\n", address)
	http.ListenAndServe(address, router)
}

// initDB opens a db connection, runs migrations, and panics if anything goes wrong
func initDB() *sql.DB {
	db, err := sql.Open("postgres", config.PostgresSourceName())
	if err != nil {
		panic(err.Error())
	}
	files, err := ioutil.ReadDir(path.Join("migrations"))
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		log.Printf("starting migration: %s\n", f.Name())
		q, err := ioutil.ReadFile(path.Join("migrations", f.Name()))
		if err != nil {
			panic(err)
		}
		if _, err = db.Exec(string(q)); err != nil {
			panic(err)
		}
		log.Printf("finished migration: %s\n", f.Name())
	}
	return db
}
