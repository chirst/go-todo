package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"user_id": 123})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

// Verifier seeks, verifies and validates JWT tokens
func Verifier(h http.Handler) http.Handler {
	return jwtauth.Verifier(tokenAuth)(h)
}

// Authenticator handles valid / invalid tokens
func Authenticator(h http.Handler) http.Handler {
	// In this example, we use
	// the provided authenticator middleware, but you can write your
	// own very easily, look at the Authenticator method in jwtauth.go
	// and tweak it, its not scary.
	return jwtauth.Authenticator(h)
}

// FromContext returns a Token and Claims
func FromContext(ctx context.Context) (*jwt.Token, jwt.MapClaims, error) {
	return jwtauth.FromContext(ctx)
}

// GetClaims gets claims
func GetClaims(ctx context.Context) jwt.MapClaims {
	_, claims, _ := jwtauth.FromContext(ctx)
	return claims
}
