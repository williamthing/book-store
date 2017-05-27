// William Thing
// 5/23/17

package main

import (
	"bookstore/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func main() {
	models.InitDB()

	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.ListenAndServe(":3000", nil)
}

// booksIndex returns all the books in the bookstore
func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	books, err := models.AllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	booksJSON, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(booksJSON)
}

// booksShow given a valid isbn will return the book with the given isbn
func booksShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	book, err := models.GetBook(isbn)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	bookJSON, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bookJSON)
}
