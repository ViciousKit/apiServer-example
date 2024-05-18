package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var errNameRequired = errors.New("name is required")
var errNameProjectIdRequired = errors.New("project id is required")
var errNameUserIdRequired = errors.New("user id is required")

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
	log.Println("handleCreateTask")
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})

		return
	}
	var task *Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})

		return
	}

	if err = validateTaskPayload(task); err != nil {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})

		return
	}

	t, err := s.store.CreateTask(task)
	if err != nil {
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating task"})

		return
	}

	WriteJson(w, http.StatusCreated, t)
}

func validateTaskPayload(task *Task) error {
	if task.Name == "" {
		return errNameRequired
	}
	if task.ProjectId == 0 {
		return errNameProjectIdRequired
	}
	if task.AssignedToId == 0 {
		return errNameUserIdRequired
	}

	return nil
}

func (s *TaskService) handleGetTask(w http.ResponseWriter, r *http.Request) {

}
