// William Thing
// 5/23/17

package main

import (
	"bookstore/models"
	"fmt"
	"log"
)

func main() {
	models.InitDB()

	books, err := models.AllBooks()
	if err != nil {
		log.Fatal(err)
	}

	for _, bk := range books {
		fmt.Printf("%s, %s, %s, $%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
