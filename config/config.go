package config

import (
	"fmt"
	"log"
	"os"
	"time"
)

// ServerAddress returns a network address to listen for requests
func ServerAddress() string {
	return fmt.Sprintf("%v:%v",
		mustDefineEnv("SERVER_ADDRESS"),
		mustDefineEnv("SERVER_PORT"),
	)
}

// JWTSignKey returns a secret key to sign JSON Web Tokens
func JWTSignKey() string {
	return mustDefineEnv("JWT_SIGN_KEY")
}

// JWTDuration returns the duration a JSON Web Token will be valid from creation
// A duration of 0 will be returned in the event of an error
func JWTDuration() time.Duration {
	d, err := time.ParseDuration(mustDefineEnv("JWT_DURATION"))
	if err != nil {
		log.Printf("failed to parse JWTDuration")
		return time.Duration(0)
	}
	return d
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
