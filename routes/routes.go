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
	app.Get("/books/:bookID", handlers.GetBookDetails)

	// Mark as Currently Reading
	app.Post("/books/:bookID/currently-reading", handlers.MarkAsCurrentlyReading)

	// Mark as Read
	app.Post("/books/:bookID/mark-as-read", handlers.MarkAsRead)

	// Mark as Did Not Finish
	app.Post("/books/:bookID/did-not-finish", handlers.MarkAsDidNotFinish)

	// Mark as To Be Read
	app.Post("/books/:bookID/to-be-read", handlers.MarkAsToBeRead)

	// Get Currently Reading List
	app.Get("/books/currently-reading-books", handlers.GetCurrentlyReadingList)

	// Get Read Books List
	app.Get("/books/read-books", handlers.GetReadBooksList)

	// Get Did Not Finish List
	app.Get("/books/did-not-finish-books", handlers.GetDidNotFinishList)

	// Get To Be Read List
	app.Get("/books/to-be-read-books", handlers.GetToBeReadList)

	// Add New Book
	app.Post("/book", handlers.AddNewBook)

	// Update Book Details
	app.Put("/books/:bookID", handlers.UpdateBookDetails)

	// Delete Book
	app.Delete("/books/:bookID", handlers.DeleteBook)

	//Add New Book View
	app.Get("/book", handlers.NewBookView)

}
