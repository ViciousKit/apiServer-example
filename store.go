package main

import (
	"database/sql"
)

type Store interface {
	CreateUser() error
}

type Repository struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (s *Repository) CreateUser() error {
	return nil
}
