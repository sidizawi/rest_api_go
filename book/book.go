package book

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	db "github.com/sidizawi/rest_api_go/database"
)

func SetupBookRoutes(router fiber.Router) {
	router.Get("/", Home)
	router.Get("/change/:id", UpdateBook)
	router.Post("/change/:id", UpdateBook)
	router.Get("/create", CreateNewBook)
	router.Post("/create", CreateNewBook)
}

func Home(c *fiber.Ctx) error {
	var books []db.Book

	db.DBConn.Find(&books)

	return c.Render("index", fiber.Map{
		"books": books,
	})
}

func UpdateBook(c *fiber.Ctx) error {
	var book db.Book

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Redirect("/")
	}

	db.DBConn.First(&book, id)

	if c.Method() == "POST" {
		if err = c.BodyParser(&book); err != nil {
			// return c.Redirect(fmt.Sprintf("/change/%v", id))
			return c.Render("change", fiber.Map{
				"book": book,
			})
		}

		db.DBConn.Save(book)
		return c.Redirect(fmt.Sprintf("/change/%v", id))
	}

	return c.Render("change", fiber.Map{
		"book": book,
	})
}

func CreateNewBook(c *fiber.Ctx) error {
	var book db.Book

	if c.Method() == "POST" {
		if err := c.BodyParser(&book); err != nil {
			// return c.Redirect("/create")
			return c.Render("create", fiber.Map{})
		}

		db.DBConn.Create(&book)
		return c.Redirect(fmt.Sprintf("/change/%v", book.ID))
	}

	return c.Render("create", fiber.Map{})
}
