package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB opens connections to mysql database given dbInfo
func InitDB(dbInfo string) {
	var err error
	db, err = sql.Open("mysql", dbInfo)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}
