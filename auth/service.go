package auth

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
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

// GetTokenForUser returns a token with the given claims
func GetTokenForUser(userID int64) (*jwt.Token, string, error) {
	return tokenAuth.Encode(jwt.MapClaims{"userID": userID})
}
