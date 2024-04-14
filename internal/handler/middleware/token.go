package middleware

import (
	"banner-service/internal/handler/tools"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var jwtKey = []byte("secret-key")

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

func TokenGen(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expirationTime := time.Now().Add(5 * time.Minute)
		claims := &Claims{
			Role: r.Header.Get("token"),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			tools.SendError(w, http.StatusInternalServerError, err.Error())
			return
		}

		r.Header.Set("Authorization", "Bearer "+tokenString)

		next.ServeHTTP(w, r)
	})
}

func WithUserAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			tools.SendStatus(w, http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len("Bearer "):]

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			tools.SendStatus(w, http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || (claims.Role != "user_token" && claims.Role != "admin_token") {
			tools.SendStatus(w, http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func WithAdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			tools.SendStatus(w, http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len("Bearer "):]

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			tools.SendStatus(w, http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || claims.Role != "admin_token" {
			tools.SendStatus(w, http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
