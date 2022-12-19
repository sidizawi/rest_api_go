package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"github.com/sidizawi/rest_api_go/book"
	"github.com/sidizawi/rest_api_go/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error

	database.DBConn, err = gorm.Open(sqlite.Open("books.db"))

	if err != nil {
		panic("Can't connect to db")
	}

	fmt.Println("Database connection succefully")

	err = database.DBConn.AutoMigrate(&book.Book{})

	if err != nil {
		panic("Can't auto migrate")
	}

	fmt.Println("Auto migrate ok!")
}

func index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func setupRoutes(app *fiber.App) {
	app.Get("/", index)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Put("/api/v1/book/:id", book.UpdateBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {

	engine := django.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	initDatabase()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
