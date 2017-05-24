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
	dbInfo := getDatabaseInfo()
	models.InitDB(dbInfo)

	books, err := models.AllBooks()
	if err != nil {
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
