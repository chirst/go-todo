package auth

// The following package is incomplete and no
// where near an acceptable solution

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	// Todo: this secret in a config
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

// Verifier is middleware for seeking, verifying and validating JWT tokens
func Verifier(h http.Handler) http.Handler {
	return jwtauth.Verifier(tokenAuth)(h)
}

// Authenticator is middleware who sends a 401 response for requests with
// bad tokens and accepts requests with good tokens
func Authenticator(h http.Handler) http.Handler {
	// Todo: think about expiration
	// Todo: this can be modified
	return jwtauth.Authenticator(h)
}

// GetUidClaim gets the userID from claims
func GetUidClaim(ctx context.Context) int64 {
	_, claims, _ := jwtauth.FromContext(ctx)
	return int64(claims["userID"].(float64))
}

// GetTokenForUser returns a token with the given claims
func GetTokenForUser(userID int64) (*jwt.Token, string, error) {
	return tokenAuth.Encode(jwt.MapClaims{"userID": userID})
}

// GenerateFromPassword returns a hashed version of the given string
func GenerateFromPassword(p string) (*string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(p), 6)
	if err != nil {
		return nil, err
	}
	sh := string(h)
	return &sh, nil
}

// CompareHashAndPassword compares a hash and a password.
// Returns nil on success, or an error on failure
func CompareHashAndPassword(h, p string) error {
	return bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
}
