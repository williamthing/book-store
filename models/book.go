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

// CreateBook will insert a book given a isbn, title, author, and price
// error if could not insert
func CreateBook(isbn string, title string, author string, price string) error {
	query, err := db.Prepare("INSERT books SET isbn=?, title=?, author=?, price=?")
	if err != nil {
		return err
	}
	result, err := query.Exec(isbn, title, author, price)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(isbn string) error {
	query, err := db.Prepare("DELETE FROM books WHERE isbn=?")
	if err != nil {
		return err
	}
	result, err := query.Exec(isbn)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
