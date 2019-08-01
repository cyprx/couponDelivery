package mysql

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "")
	if err != nil {
		log.Fatalf("Could not open database connection!")
	}

	DB = db
}

func GetDB() *sql.DB {
	return DB
}
