package auth

// The following package is incomplete and should not be trusted

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chirst/go-todo/config"
	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwt"

	"golang.org/x/crypto/bcrypt"
)

type uidType string

const UIDKey uidType = "uid"

var tokenAuth *jwtauth.JWTAuth

func init() {
	key := config.JWTSignKey()
	tokenAuth = jwtauth.New("HS256", []byte(key), nil)
}

// Verifier is middleware for seeking, verifying and validating JWT tokens.
func Verifier(h http.Handler) http.Handler {
	return jwtauth.Verifier(tokenAuth)(h)
}

// Authenticator is middleware who sends a 401 response for requests with bad
// tokens and accepts requests with good tokens.
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// The following implementation comes from jwtauth.Authenticator, but is
		// enhanced to check for expired tokens.
		token, claims, err := jwtauth.FromContext(r.Context())

		// Checks from the jwt.Authenticator to see if the token is valid
		if err != nil {
			writeUnauthorized(&w, err.Error())
			return
		}
		if token == nil {
			writeUnauthorized(&w, "token is nil")
			return
		}
		if err = jwt.Validate(token); err != nil {
			writeUnauthorized(&w, err.Error())
			return
		}

		// Check if token is expired.
		expiry, err := getExpiry(claims)
		if err != nil {
			writeUnauthorized(&w, err.Error())
			return
		}
		if time.Now().Unix() > expiry {
			writeUnauthorized(&w, "token expired")
			return
		}

		uid, err := getUIDClaim(r.Context())
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to get user", http.StatusUnauthorized)
			return
		}
		if uid == nil {
			http.Error(w, "Unable to get user", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UIDKey, *uid)
		r2 := r.WithContext(ctx)

		next.ServeHTTP(w, r2)
	})
}

func writeUnauthorized(w *http.ResponseWriter, message string) {
	log.Print(message)
	http.Error(
		*w,
		http.StatusText(http.StatusUnauthorized),
		http.StatusUnauthorized,
	)
}

func getExpiry(claims map[string]interface{}) (int64, error) {
	e, ok := claims["expires"].(float64)
	if !ok {
		return 0, errors.New("unable to parse expires from claims")
	}
	return int64(e), nil
}

func getUIDClaim(ctx context.Context) (*int, error) {
	_, claims, _ := jwtauth.FromContext(ctx)
	if t, ok := claims["userID"].(float64); ok {
		ti := int(t)
		return &ti, nil
	}
	if t, ok := claims["userID"].(int); ok {
		return &t, nil
	}
	return nil, fmt.Errorf("failed to assert type of %v", claims["userID"])
}

// GetTokenForUser returns a token with the given claims.
func GetTokenForUser(userID int) (jwt.Token, string, error) {
	return tokenAuth.Encode(map[string]interface{}{
		"userID":  userID,
		"expires": time.Now().Add(config.JWTDuration()).Unix(),
	})
}

// GenerateFromPassword returns a hashed version of the given string.
func GenerateFromPassword(p string) (*string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(p), 6)
	if err != nil {
		return nil, err
	}
	sh := string(h)
	return &sh, nil
}

// CompareHashAndPassword compares a hash and a password returning an error when
// the hash an password do not match.
func CompareHashAndPassword(h, p string) error {
	return bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
}
