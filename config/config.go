package config

import (
	"fmt"
	"log"
	"os"
	"time"
)

// JWTSignKey returns a secret key to sign JSON Web Tokens
func JWTSignKey() string {
	return getEnv("JWT_SIGN_KEY", "secret")
}

// JWTDuration returns the duration a JSON Web Token will be valid from creation
// A duration of 0 will be returned in the event of an error
func JWTDuration() time.Duration {
	d, err := time.ParseDuration(getEnv("JWT_DURATION", "24h"))
	if err != nil {
		log.Printf("failed to parse JWTDuration")
		return time.Duration(0)
	}
	return d
}

// ServerAddress returns a network address to listen for requests
func ServerAddress() string {
	return fmt.Sprintf("%v:%v",
		getEnv("SERVER_ADDRESS", "0.0.0.0"),
		getEnv("SERVER_PORT", "3000"),
	)
}

// PostgresSourceName returns a string containing connection info specific to a
// Postgres database
func PostgresSourceName() string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		getEnv("POSTGRES_USER", "postgres"),
		getEnv("POSTGRES_PASSWORD", "12345"),
		getEnv("POSTGRES_HOST", "127.0.0.1"),
		getEnv("POSTGRES_PORT", "5432"),
		getEnv("POSTGRES_DB", "todo"),
	)
}

// SkipPostgres is true if tests requiring postgres should be skipped
func SkipPostgres() bool {
	r := getEnv("POSTGRES_TEST", "")
	return r == ""
}

func getEnv(key string, defaultVal string) string {
	v, found := os.LookupEnv(key)
	if !found {
		return defaultVal
	}
	return v
}
