package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type TaskService struct {
	store Store
}

func NewTaskService(s Store) *TaskService {
	return &TaskService{store: s}
}

func (s *TaskService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", s.handleCreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", s.handleGetTask).Methods("GET")
}

func (s *TaskService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return
	}
	var task *Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		return
	}

	if err = validateTaskPayload(task); err != nil {
		return
	}
}

func validateTaskPayload(task *Task) error {

}

func (s *TaskService) handleGetTask(w http.ResponseWriter, r *http.Request) {

}
