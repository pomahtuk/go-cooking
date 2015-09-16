package books

import (
	"github.com/go-martini/martini"
)

func GetBooks(params martini.Params) string {
	return "Got books!"
}

func NewBook(params martini.Params) string {
	return "New book!"
}

func UpdateBook(params martini.Params) string {
	return "Updated book!"
}

func DeleteBook(params martini.Params) string {
	return "Deleted book!"
}
