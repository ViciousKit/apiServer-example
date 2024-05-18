package main

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type Task struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	Status       string    `json:"status"`
	ProjectId    int64     `json:"projectId"`
	AssignedToId int64     `json:"assignedToId"`
	Created      time.Time `json:"created"`
}

type Project struct {
	Id      int64     `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

type User struct {
	Id           int64     `json:"id"`
	Email        string    `json:"email"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Status       string    `json:"status"`
	ProjectId    int64     `json:"projectId"`
	AssignedToId int64     `json:"assignedToId"`
	Created      time.Time `json:"created"`
}
