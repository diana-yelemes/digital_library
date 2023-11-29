package handlers

import (
	"fmt"

	"github.com/diana-yelemes/digital_library/database"
	"github.com/diana-yelemes/digital_library/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllBooks(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Db.Find(&books)

	return c.Render("index", fiber.Map{
		"Title":    "my library",
		"Subtitle": "all of the books i own: ",
		"Books":    books,
	})
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
	// Get bookID from params
	bookID := c.Params("bookID")

	// Retrieve the book from the database
	var book models.Book
	result := database.DB.Db.First(&book, bookID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	// Update the book status
	book.Status = status
	database.DB.Db.Save(&book)

	// Return the updated book
	return c.JSON(book)
}

func GetBooksByStatus(c *fiber.Ctx, status string) error {
	fmt.Println("Status:", status)

	var books []models.Book

	// Construct the SQL query dynamically using the status parameter
	sql := "SELECT * FROM books WHERE status = '" + status + "' AND deleted_at IS NULL ORDER BY id"

	// Execute the dynamically constructed SQL query
	err := database.DB.Db.Raw(sql).Find(&books).Error
	if err != nil {
		fmt.Println("Error fetching books by status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(books)
}

func GetToBeReadBooks(c *fiber.Ctx) error {
	var books []models.Book

	sql := "SELECT * FROM books WHERE status = $1 AND deleted_at IS NULL ORDER BY id"
	err := database.DB.Db.Raw(sql, "to-be-read").Find(&books).Error

	if err != nil {
		fmt.Println("Error fetching books by status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(books)
}

func GetCurrentlyReadingBooks(c *fiber.Ctx) error {
	var books []models.Book

	sql := "SELECT * FROM books WHERE status = $1 AND deleted_at IS NULL ORDER BY id"
	err := database.DB.Db.Raw(sql, "currently-reading").Find(&books).Error

	if err != nil {
		fmt.Println("Error fetching books by status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(books)
}

func GetReadBooks(c *fiber.Ctx) error {
	var books []models.Book

	sql := "SELECT * FROM books WHERE status = $1 AND deleted_at IS NULL ORDER BY id"
	err := database.DB.Db.Raw(sql, "read").Find(&books).Error

	if err != nil {
		fmt.Println("Error fetching books by status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(books)
}

func GetDidNotFinishBooks(c *fiber.Ctx) error {
	fmt.Println("Status:", "did-not-finish")

	var books []models.Book
	sql := "SELECT * FROM books WHERE status = ? AND books.deleted_at IS NULL ORDER BY id"

	err := database.DB.Db.Raw(sql, "did-not-finish").Find(&books).Error
	if err != nil {
		fmt.Println("Error fetching books by status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(books)
}
func NewBookView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Book",
		"Subtitle": "Create an addition to your library: ",
	})
}

func AddNewBook(c *fiber.Ctx) error {
	var newBook models.Book
	if err := c.BodyParser(&newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	database.DB.Db.Create(&newBook)
	return ConfirmationView(c)
}

func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title":    "Book added successfully",
		"Subtitle": "You can add more books to your library",
	})
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
	return c.JSON(updatedBook)
}

func DeleteBook(c *fiber.Ctx) error {
	bookID := c.Params("bookID")
	result := database.DB.Db.Delete(&models.Book{}, bookID)
	if result.Error != nil || result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	c.SendStatus(fiber.StatusNoContent)
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book was deleted"})
}
