package main

import (
	"./controllers/books"
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()

	m.Use(martini.Logger())
	m.Use(martini.Recovery())
	m.Use(martini.Static("public"))

	m.Get("/", func() string {
		return "Hello world!"
	})

	// m.Get("/hello/(?P<name>[a-zA-Z]+)", func(params martini.Params) string {
	// 	return fmt.Sprintf("Hello %s", params["name"])
	// })

	m.Group("/books", func(r martini.Router) {
		r.Get("/:id", books.GetBooks)
		r.Post("/new", books.NewBook)
		r.Put("/update/:id", books.UpdateBook)
		r.Delete("/delete/:id", books.DeleteBook)
	})

	m.Run()
}
