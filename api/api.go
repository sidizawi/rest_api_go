package api

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/sidizawi/rest_api_go/database"
)

func SetupApiRoutes(router fiber.Router) {
	router.Get("/book", GetBooks)
	router.Get("/book/:id", GetBook)
	router.Post("/book", NewBook)
	router.Put("/book/:id", UpdateBook)
	router.Delete("/book/:id", DeleteBook)
}

func GetBooks(c *fiber.Ctx) error {
	var books []db.Book

	db.DBConn.Find(&books)

	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	var book db.Book
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.JSON(map[string]string{"error": "id param is not an int"})
	}

	db.DBConn.First(&book, id)

	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	var book db.Book

	if err := c.BodyParser(&book); err != nil {
		return c.JSON(map[string]string{"error": "can't parse the body"})
	}

	db.DBConn.Create(&book)

	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	var book db.Book
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.JSON(map[string]string{"error": "id param is not an int"})
	}

	db.DBConn.First(&book, id)

	if err = c.BodyParser(&book); err != nil {
		return c.JSON(map[string]string{"error": "can't parse the body"})
	}

	db.DBConn.Save(book)

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.JSON(map[string]string{"error": "id param is not an int"})
	}

	db.DBConn.Delete(&db.Book{}, id)

	return c.JSON(map[string]string{"message": "book deleted"})
}
