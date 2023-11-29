// routes/routes.go

package routes

import (
	"github.com/diana-yelemes/digital_library/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Get All Books
	app.Get("/", handlers.GetAllBooks)

	// Get Book Details
	app.Get("/api/books/:bookID", handlers.GetBookDetails)

	// Mark as Currently Reading
	app.Post("/api/books/:bookID/currently-reading", handlers.MarkAsCurrentlyReading)

	// Mark as Read
	app.Post("/api/books/:bookID/read", handlers.MarkAsRead)

	// Mark as Did Not Finish
	app.Post("/api/books/:bookID/did-not-finish", handlers.MarkAsDidNotFinish)

	// Mark as To Be Read
	app.Post("/api/books/:bookID/to-be-read", handlers.MarkAsToBeRead)

	// Get Currently Reading List
	app.Get("/api/currently-reading", handlers.GetCurrentlyReadingBooks)

	// Get Read Books List
	app.Get("/api/read", handlers.GetReadBooks)

	// Get Did Not Finish List
	app.Get("/api/did-not-finish", handlers.GetDidNotFinishBooks)

	// Get To Read List
	app.Get("/api/to-be-read", handlers.GetToBeReadBooks)

	// Add New Book
	app.Post("/book", handlers.AddNewBook)

	// Update Book Details
	app.Put("/api/books/:bookID", handlers.UpdateBookDetails)

	// Delete Book
	app.Delete("/api/books/:bookID", handlers.DeleteBook)

	//Add New Book View
	app.Get("/book", handlers.NewBookView)

}
