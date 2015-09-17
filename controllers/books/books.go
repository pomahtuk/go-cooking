package books

import (
	"github.com/go-martini/martini"
)

// GetBooks will return book for given id
func GetBooks(params martini.Params) string {
	return "Got books!"
}

// NewBook will create a new book
func NewBook(params martini.Params) string {
	return "New book!"
}

// UpdateBook will update book
func UpdateBook(params martini.Params) string {
	return "Updated book!"
}

// DeleteBook wiil delete book
func DeleteBook(params martini.Params) string {
	return "Deleted book!"
}
