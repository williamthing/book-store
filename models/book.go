package models

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

// AllBooks returns all books in the bookstore
func AllBooks() ([]*Book, error) {
	var books []*Book

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := new(Book)
		err := rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

// GetBook will return the book with the given isbn if found
// error if not
func GetBook(isbn string) (*Book, error) {
	sqlQuery := "SELECT * FROM books WHERE isbn = " + isbn
	row := db.QueryRow(sqlQuery)
	book := new(Book)
	err := row.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)

	return book, err
}
