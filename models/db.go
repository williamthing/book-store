package models

import (
	"bufio"
	"bytes"
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB opens connections to mysql database given dbInfo
func InitDB() {
	dbInfo := getDatabaseInfo()
	var err error
	db, err = sql.Open("mysql", dbInfo)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

// getDatabaseInfo returns database connection information required: username,
// password, and database endpoint
func getDatabaseInfo() string {
	var dbInfo []string
	var buffer bytes.Buffer

	file, err := os.Open("models/dbinfo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		info := scanner.Text()
		dbInfo = append(dbInfo, info)
	}

	user := dbInfo[0]
	pw := dbInfo[1]
	dbConnection := dbInfo[2]

	buffer.WriteString(user)
	buffer.WriteString(":")
	buffer.WriteString(pw)
	buffer.WriteString("@tcp(")
	buffer.WriteString(dbConnection)
	buffer.WriteString(":3306)/bookstore")
	param := buffer.String()
	return param
}
