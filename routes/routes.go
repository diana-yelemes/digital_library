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
	app.Post("/api/books/:bookID/mark-as-read", handlers.MarkAsRead)

	// Mark as Did Not Finish
	app.Post("/api/books/:bookID/did-not-finish", handlers.MarkAsDidNotFinish)

	// Mark as To Be Read
	app.Post("/api/books/:bookID/to-be-read", handlers.MarkAsToBeRead)

	// Get Currently Reading List
	app.Get("/api/books/currently-reading-books", handlers.GetCurrentlyReadingList)

	// Get Read Books List
	app.Get("/api/books/read-books", handlers.GetReadBooksList)

	// Get Did Not Finish List
	app.Get("/api/books/did-not-finish-books", handlers.GetDidNotFinishList)

	// Get To Be Read List
	app.Get("/api/books/to-be-read-books", handlers.GetToBeReadList)

	// Add New Book
	app.Post("/book", handlers.AddNewBook)

	// Update Book Details
	app.Put("/api/books/:bookID", handlers.UpdateBookDetails)

	// Delete Book
	app.Delete("/api/books/:bookID", handlers.DeleteBook)

	//Add New Book View
	app.Get("/book", handlers.NewBookView)

}
