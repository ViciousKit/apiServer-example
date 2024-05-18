package main

import (
	"database/sql"
)

type Store interface {
	CreateUser(t *User) (*User, error)
	CreateTask(t *Task) (*Task, error)
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
