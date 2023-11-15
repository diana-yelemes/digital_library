// cmd/routes.go

package main

import (
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Div Rhino Trivia App!")
	})
	// Get All Books
	app.Get("/api/books", getAllBooks)

	// Get Book Details
	app.Get("/api/books/:bookID", getBookDetails)

	// Mark as Currently Reading
	app.Post("/api/books/:bookID/currently-reading", markAsCurrentlyReading)

	// Mark as Read
	app.Post("/api/books/:bookID/mark-as-read", markAsRead)

	// Mark as Did Not Finish
	app.Post("/api/books/:bookID/did-not-finish", markAsDidNotFinish)

	// Get Currently Reading List
	app.Get("/api/currently-reading", getCurrentlyReadingList)

	// Get Read Books List
	app.Get("/api/read-books", getReadBooksList)

	// Get Did Not Finish List
	app.Get("/api/did-not-finish-books", getDidNotFinishList)

	// Add New Book
	app.Post("/api/books", addNewBook)

	// Update Book Details
	app.Put("/api/books/:bookID", updateBookDetails)

	// Delete Book
	app.Delete("/api/books/:bookID", deleteBook)

	// Search Books
	app.Get("/api/books/search", searchBooks)

	// Mark as To Be Read
	app.Post("/api/books/:bookID/to-be-read", markAsToBeRead)

	// Get To Be Read List
	app.Get("/api/to-be-read", getToBeReadList)

}
