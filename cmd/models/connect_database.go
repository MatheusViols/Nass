package models

import (
	"log"

	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	cfg := mysql.NewConfig()
	cfg.User = "nassuser"
	cfg.Passwd = "nassuser"
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "nass"

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(err)
	}

	log.Print("Successfuly connected to the database")

	return db
}