package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MysqlStorage struct {
	db *sql.DB
}

func NewMysqlStorage(cfg mysql.Config) *MysqlStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connections success")

	return &MysqlStorage{db}
}

func (s *MysqlStorage) Init() (*sql.DB, error) {
	//initialize tables
	if err := s.createUserTable(); err != nil {
		return nil, err
	}
	if err := s.createProjectTable(); err != nil {
		return nil, err
	}
	if err := s.createTaskTable(); err != nil {
		return nil, err
	}

	return s.db, nil
}

func (s *MysqlStorage) createProjectTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS project (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		) ENGINE=InnoDB CHARACTER SET utf8;
	`)
	if err != nil {
		return err
	}

	return nil
}

func (s *MysqlStorage) createTaskTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS task (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			status ENUM('TODO', 'IN_PROGRESS', 'IN_TESTING', 'DONE') NOT NULL DEFAULT 'TODO',
			project_id INT UNSIGNED NOT NULL,
			assigned_to_id INT UNSIGNED NOT NULL,
			created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			FOREIGN KEY (assigned_to_id) REFERENCES user(id),
			FOREIGN KEY (project_id) REFERENCES project(id)
		) ENGINE=InnoDB CHARACTER SET utf8;
	`)
	if err != nil {
		return err
	}

	return nil
}

func (s *MysqlStorage) createUserTable() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS user (
			id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			email VARCHAR(255) NOT NULL,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id),
			UNIQUE KEY (email)
		) ENGINE=InnoDB CHARACTER SET utf8;
	`)
	if err != nil {
		return err
	}

	return nil
}
