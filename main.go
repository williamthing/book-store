// William Thing
// 5/23/17

package main

import (
	"bookstore/models"
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	var books []*models.Book

	dbInfo := getDatabaseInfo()
	models.InitDB(dbInfo)

	db := models.InitDB(dbInfo)
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		book := new(models.Book)
		err := rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, bk := range books {
		fmt.Printf("%s, %s, %s, $%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}

//	returns database connection information required: username, password,
//	and database endpoint
func getDatabaseInfo() string {
	var dbInfo []string
	var buffer bytes.Buffer

	file, err := os.Open("dbinfo.txt")
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
