package book

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sidizawi/rest_api_go/database"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string
	Author string
	Rating int
}

func GetBooks(c *fiber.Ctx) error {
	var books []Book

	result := database.DBConn.Find(&books)

	if result.Error != nil {
		return c.Render("index", fiber.Map{
			"books": []Book{},
		}, "books/index")
	}

	return c.Render("index", fiber.Map{
		"books": books,
	}, "books/index")
}

func GetBook(c *fiber.Ctx) error {
	var book Book
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Render("index", fiber.Map{
			"can_see": false,
			// "book":    Book{},
		}, "books/book")
	}

	database.DBConn.First(&book, id)

	return c.Render("index", fiber.Map{
		"can_see": true,
		"book":    book,
	}, "books/book")
}

func NewBook(c *fiber.Ctx) error {
	var book Book

	if err := c.BodyParser(&book); err != nil {
		c.Path("/")
		return c.RestartRouting()
	}

	database.DBConn.Create(&book)

	c.Path("/api/v1/book")
	return c.RestartRouting()
}

func UpdateBook(c *fiber.Ctx) error {
	var book Book
	id, err := c.ParamsInt("id")

	if err != nil {
		c.Path("/")
		return c.RestartRouting()
	}

	database.DBConn.First(&book, id)

	if err = c.BodyParser(&book); err != nil {
		c.Path("/")
		return c.RestartRouting()
	}

	database.DBConn.Save(book)

	c.Path(fmt.Sprintf("/api/v1/book/%v", book.ID))
	return c.RestartRouting()
}

func DeleteBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		c.Path("/")
		return c.RestartRouting()
	}

	database.DBConn.Delete(&Book{}, id)

	c.Path("/api/v1/book")
	return c.RestartRouting()
}
