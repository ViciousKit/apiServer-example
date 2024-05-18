package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 "test",
		Passwd:               "secret",
		DBName:               "projectManager",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	sqlStorage := NewMysqlStorage(cfg)

	db, err := sqlStorage.Init()

	if err != nil {
		log.Fatal(err)
	}
	store := NewStore(db)

	api := NewApiServer(":8081", store)
	api.Serve()
}
