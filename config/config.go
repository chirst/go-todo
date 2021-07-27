package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// InitConfig reads in config from a file or panics on failure
func InitConfig() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("fatal error in config: %s", err))
	}
}

// ServerAddress returns a network address to listen for requests
func ServerAddress() string {
	return fmt.Sprintf("%v:%v",
		mustDefineEnv("SERVER_ADDRESS"),
		mustDefineEnv("SERVER_PORT"),
	)
}

// UseMemoryDB returns true when the use memory db config can be found and is
// set to true otherwise UseMemoryDB will return false.
func UseMemoryDB() bool {
	return mustDefineEnv("USE_MEMORY_DB") == "true"
}

// PostgresSourceName returns a string containing connection info specific to a
// Postgres database
func PostgresSourceName() string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		mustDefineEnv("POSTGRES_USER"),
		mustDefineEnv("POSTGRES_PASSWORD"),
		mustDefineEnv("POSTGRES_HOST"),
		mustDefineEnv("POSTGRES_PORT"),
		mustDefineEnv("POSTGRES_DB"),
	)
}

func mustDefineEnv(key string) string {
	v, found := os.LookupEnv(key)
	if !found {
		log.Panicf("environment variable with key %v is not found", key)
	}
	return v
}
