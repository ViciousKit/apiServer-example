package main

import (
	"database/sql"
)

type Store interface {
	CreateUser(t *User) (*User, error)
	CreateTask(t *Task) (*Task, error)
	GetTask(id string) (*Task, error)
	GetUserById(id string) (*User, error)
}

type Repository struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (s *Repository) CreateUser(u *User) (*User, error) {
	return nil, nil
}

func (s *Repository) CreateTask(t *Task) (*Task, error) {
	rows, err := s.db.Exec("INSERT INTO task (name, status, project_id, assigned_to) VALUES (?, ?, ?, ?)", t.Name, t.Status, t.ProjectId, t.AssignedToId)
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	t.Id = id

	return t, nil
}

func (s *Repository) GetTask(id string) (*Task, error) {
	var t Task
	err := s.db.QueryRow("SELECT id, name, status, project_id, assigned_to, created FROM task WHERE id = ?", id).Scan(&t.Id, &t.Name, &t.Status, &t.ProjectId, &t.AssignedToId, &t.Created)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *Repository) GetUserById(id string) (*User, error) {
	var u User
	err := s.db.QueryRow("SELECT id, email, firstName, lastName, created FROM user WHERE id = ?", id).Scan(&u.Id, &u.FirstName, &u.LastName, &u.Created)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
