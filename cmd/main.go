package main

import (
	"github.com/diana-yelemes/digital_library/database"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

// Book model
type Book struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Status      string `json:"status"` // "read", "to read", "currently reading", "did not finish"
}

func getAllBooks(c *fiber.Ctx) {
	// Implement logic to retrieve all books from the database
	// and return as JSON
}

func getBookDetails(c *fiber.Ctx) {
	// Implement logic to retrieve book details based on bookID
	// from the database and return as JSON
}

func markAsCurrentlyReading(c *fiber.Ctx) {
	// Implement logic to mark a book as "Currently Reading" in the database
}

func markAsRead(c *fiber.Ctx) {
	// Implement logic to mark a book as "Read" in the database
}

func markAsDidNotFinish(c *fiber.Ctx) {
	// Implement logic to mark a book as "Did Not Finish" in the database
}

func getCurrentlyReadingList(c *fiber.Ctx) {
	// Implement logic to retrieve the list of books marked as "Currently Reading"
	// from the database and return as JSON
}

func getReadBooksList(c *fiber.Ctx) {
	// Implement logic to retrieve the list of books marked as "Read"
	// from the database and return as JSON
}

func getDidNotFinishList(c *fiber.Ctx) {
	// Implement logic to retrieve the list of books marked as "Did Not Finish"
	// from the database and return as JSON
}

func addNewBook(c *fiber.Ctx) {
	// Implement logic to add a new book to the database
}

func updateBookDetails(c *fiber.Ctx) {
	// Implement logic to update book details based on bookID in the database
}

func deleteBook(c *fiber.Ctx) {
	// Implement logic to delete a book based on bookID from the database
}

func searchBooks(c *fiber.Ctx) {
	// Implement logic to search for books based on a query parameter
	// from the database and return as JSON
}

func markAsToBeRead(c *fiber.Ctx) {
	// Implement logic to mark a book as "To Be Read" in the database
}

func getToBeReadList(c *fiber.Ctx) {
	// Implement logic to retrieve the list of books marked as "To Be Read"
	// from the database and return as JSON
}

func main() {
	// Initialize Fiber app and database connection
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	// Start the Fiber app
	app.Listen(3000)
}
