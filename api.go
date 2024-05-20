package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	address string
	store   Store
}

func NewApiServer(addr string, store Store) *ApiServer {
	return &ApiServer{
		address: addr,
		store:   store,
	}
}

func (s *ApiServer) Serve() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.Methods()

	taskService := NewTaskService(s.store)
	taskService.RegisterRoutes(subRouter)

	userService := NewUserService(s.store)
	userService.RegisterRoutes(subRouter)

	log.Println("Starting api server at", s.address)

	log.Fatal(http.ListenAndServe(s.address, subRouter))
}
