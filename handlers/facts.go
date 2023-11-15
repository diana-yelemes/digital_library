package handlers

import (
	"github.com/diana-yelemes/digital_library/database"
	"github.com/diana-yelemes/digital_library/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllBooks(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Db.Find(&books)
	return c.JSON(books)
}

func GetBookDetails(c *fiber.Ctx) error {
	bookID := c.Params("bookID")
	var book models.Book
	result := database.DB.Db.First(&book, bookID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

func MarkAsCurrentlyReading(c *fiber.Ctx) error {
	return markBookStatus(c, "currently reading")
}

func MarkAsRead(c *fiber.Ctx) error {
	return markBookStatus(c, "read")
}

func MarkAsDidNotFinish(c *fiber.Ctx) error {
	return markBookStatus(c, "did not finish")
}

func MarkAsToBeRead(c *fiber.Ctx) error {
	return markBookStatus(c, "to read")
}

func markBookStatus(c *fiber.Ctx, status string) error {
	// Implement logic to mark a book with the specified status in the database
	return nil
}

func getBooksByStatus(c *fiber.Ctx, status string) error {
	var books []models.Book
	database.DB.Db.Where("status = ?", status).Find(&books)
	return c.JSON(books)
}

func GetCurrentlyReadingList(c *fiber.Ctx) error {
	return getBooksByStatus(c, "currently reading")
}

func GetReadBooksList(c *fiber.Ctx) error {
	return getBooksByStatus(c, "read")
}

func GetDidNotFinishList(c *fiber.Ctx) error {
	return getBooksByStatus(c, "did not finish")
}

func GetToBeReadList(c *fiber.Ctx) error {
	return getBooksByStatus(c, "to read")
}

func AddNewBook(c *fiber.Ctx) error {
	var newBook models.Book
	if err := c.BodyParser(&newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	database.DB.Db.Create(&newBook)
	return c.JSON(newBook)
}

func UpdateBookDetails(c *fiber.Ctx) error {
	bookID := c.Params("bookID")
	var updatedBook models.Book
	if err := c.BodyParser(&updatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	result := database.DB.Db.Model(&models.Book{}).Where("id = ?", bookID).Updates(&updatedBook)
	if result.Error != nil || result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	c.SendStatus(fiber.StatusNoContent)
	return nil
}

func DeleteBook(c *fiber.Ctx) error {
	bookID := c.Params("bookID")
	result := database.DB.Db.Delete(&models.Book{}, bookID)
	if result.Error != nil || result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	c.SendStatus(fiber.StatusNoContent)
	return nil
}

func SearchBooks(c *fiber.Ctx) error {
	query := c.Query("q")
	var searchResults []models.Book
	database.DB.Db.Where("title LIKE ? OR author LIKE ?", "%"+query+"%", "%"+query+"%").Find(&searchResults)
	return c.JSON(searchResults)
}
