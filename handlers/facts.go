//handlers/facts.go

package handlers

import (
	"github.com/diana-yelemes/digital_library/database"
	"github.com/diana-yelemes/digital_library/models"
	"github.com/gofiber/fiber"
)

func getAllBooks(c *fiber.Ctx) {
	var books []models.Book
	database.DB.Db.Find(&books)
	c.JSON(books)
}

func getBookDetails(c *fiber.Ctx) {
	bookID := c.Params("bookID")
	var book models.Book
	result := database.DB.Db.First(&book, bookID)
	if result.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		return
	}
	c.JSON(book)
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

func markAsToBeRead(c *fiber.Ctx) {
	// Implement logic to mark a book as "To Be Read" in the database
}

func getCurrentlyReadingList(c *fiber.Ctx) {
	var currentlyReadingBooks []models.Book
	database.DB.Db.Where("status = ?", "currently reading").Find(&currentlyReadingBooks)
	c.JSON(currentlyReadingBooks)
}

func getReadBooksList(c *fiber.Ctx) {
	var readBooks []models.Book
	database.DB.Db.Where("status = ?", "read").Find(&readBooks)
	c.JSON(readBooks)
}

func getDidNotFinishList(c *fiber.Ctx) {
	var didNotFinishBooks []models.Book
	database.DB.Db.Where("status = ?", "did not finish").Find(&didNotFinishBooks)
	c.JSON(didNotFinishBooks)
}

func getToBeReadList(c *fiber.Ctx) {
	var toBeReadBooks []models.Book
	database.DB.Db.Where("status = ?", "to read").Find(&toBeReadBooks)
	c.JSON(toBeReadBooks)
}

func addNewBook(c *fiber.Ctx) {
	var newBook models.Book
	if err := c.BodyParser(&newBook); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		return
	}

	database.DB.Db.Create(&newBook)
	c.JSON(newBook)
}

func updateBookDetails(c *fiber.Ctx) {
	bookID := c.Params("bookID")
	var updatedBook models.Book
	if err := c.BodyParser(&updatedBook); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
		return
	}

	result := database.DB.Db.Model(&models.Book{}).Where("id = ?", bookID).Updates(&updatedBook)
	if result.Error != nil || result.RowsAffected == 0 {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		return
	}

	c.SendStatus(fiber.StatusNoContent)
}

func deleteBook(c *fiber.Ctx) {
	bookID := c.Params("bookID")
	result := database.DB.Db.Delete(&models.Book{}, bookID)
	if result.Error != nil || result.RowsAffected == 0 {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		return
	}

	c.SendStatus(fiber.StatusNoContent)
}

func searchBooks(c *fiber.Ctx) {
	query := c.Query("q")
	var searchResults []models.Book
	database.DB.Db.Where("title LIKE ? OR author LIKE ?", "%"+query+"%", "%"+query+"%").Find(&searchResults)
	c.JSON(searchResults)
}
