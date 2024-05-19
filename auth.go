package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func WithJWTAuth(handlerFunc http.HandlerFunc, store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := GetTokenFromRequest(r)
		token, err := validateToken(tokenString)
		if err != nil || !token.Valid {
			log.Println("failed to auth token")
			WriteJson(w, http.StatusUnauthorized, ErrorResponse{Error: fmt.Errorf("permission denied").Error()})

			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId := claims["userID"].(string)

		_, err = store.GetUserById(userId)
		if err != nil {
			log.Println("user not found")
			WriteJson(w, http.StatusUnauthorized, ErrorResponse{Error: fmt.Errorf("user not found").Error()})

			return
		}
		handlerFunc(w, r)
	}
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}

func validateToken(token string) (*jwt.Token, error) {
	secret := Envs.JWTSecret

	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secret), nil
	})
}
