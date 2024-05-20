package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserService struct {
	store Store
}

func NewUserService(s Store) *UserService {
	return &UserService{store: s}
}

var errFirstNRequired = errors.New("first name required")
var errLastNRequired = errors.New("last name required")
var errEmailRequired = errors.New("email required")
var errPasswordRequired = errors.New("password required")

func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/user", s.handleRegisterUser).Methods("POST")
}

func (s *UserService) handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})

		return
	}

	var user *User
	err = json.Unmarshal(body, &user)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})

		return
	}

	if err := validateUserPayload(user); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})

		return
	}

	pwHash, err := HashPassword(user.Password)
	if err != nil {
		log.Println(err)

		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Failed to create user"})

		return
	}
	user.Password = pwHash

	u, err := s.store.CreateUser(user)
	if err != nil {
		log.Println(err)
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Failed to create user"})

		return
	}
	token, err := createAndSetAuthCookie(u.Id, w)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating session"})

		return
	}

	WriteJson(w, http.StatusOK, token)
}

func createAndSetAuthCookie(id int64, w http.ResponseWriter) (string, error) {
	secret := []byte(Envs.JWTSecret)
	token, err := CreateJWT(secret, id)
	if err != nil {
		return "", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})

	return token, nil
}

func validateUserPayload(u *User) error {
	if u.FirstName == "" {
		return errFirstNRequired
	}
	if u.LastName == "" {
		return errLastNRequired
	}
	if u.Email == "" {
		return errEmailRequired
	}
	if u.Password == "" {
		return errPasswordRequired
	}

	return nil
}
