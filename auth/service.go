package auth

// The following package is incomplete and should not be trusted

import (
	"context"
	"net/http"
	"time"

	"github.com/chirst/go-todo/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	key := config.JWTSignKey()
	tokenAuth = jwtauth.New("HS256", []byte(key), nil)
}

// Verifier is middleware for seeking, verifying and validating JWT tokens
func Verifier(h http.Handler) http.Handler {
	return jwtauth.Verifier(tokenAuth)(h)
}

// Authenticator is middleware who sends a 401 response for requests with bad
// tokens and accepts requests with good tokens. This implementation comes from
// jwtauth.Authenticator, but is enhanced to check for expired tokens.
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())

		// Checks from the jwt.Authenticator to see if the token is valid
		if err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}
		if token == nil || !token.Valid {
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}

		// Check if token is expired
		e, ok := claims["expires"].(float64)
		if !ok {
			http.Error(
				w,
				"unauthorized failed to parse expiration",
				http.StatusUnauthorized,
			)
			return
		}
		if int64(e) < time.Now().Unix() {
			http.Error(w, "unauthorized token expired", http.StatusUnauthorized)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

// GetUIDClaim gets the userID from claims
func GetUIDClaim(ctx context.Context) int64 {
	_, claims, _ := jwtauth.FromContext(ctx)
	if t, ok := claims["userID"].(float64); ok {
		return int64(t)
	}
	return claims["userID"].(int64)
}

// GetTokenForUser returns a token with the given claims
func GetTokenForUser(userID int64) (*jwt.Token, string, error) {
	return tokenAuth.Encode(jwt.MapClaims{
		"userID":  userID,
		"expires": time.Now().Add(config.JWTDuration()).Unix(),
	})
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

// CompareHashAndPassword compares a hash and a password returning an error when
// the hash an password do not match
func CompareHashAndPassword(h, p string) error {
	return bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
}
