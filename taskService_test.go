package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateTask(t *testing.T) {
	ms := &MockStore{}
	service := NewTaskService(ms)

	t.Run("should return error if name is empty", func(t *testing.T) {
		payload := &Task{
			Name: "",
		}
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/task", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/task", service.handleCreateTask).Methods("POST")

		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusBadRequest {
			t.Error("invalid status code, shoud be bad request")
		}
	})

	t.Run("should create task", func(t *testing.T) {
		payload := &Task{
			Name:         "Test",
			ProjectId:    2,
			AssignedToId: 2,
		}
		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/task", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/task", service.handleCreateTask).Methods("POST")

		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusCreated {
			t.Errorf("expected status code %d, but got %d", http.StatusCreated, recorder.Code)
		}
	})
}
