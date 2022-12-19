package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"github.com/sidizawi/rest_api_go/api"
	"github.com/sidizawi/rest_api_go/book"
	db "github.com/sidizawi/rest_api_go/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error

	db.DBConn, err = gorm.Open(sqlite.Open("books.db"))

	if err != nil {
		panic("Can't connect to db")
	}

	fmt.Println("Database connection succefully")

	err = db.DBConn.AutoMigrate(&db.Book{})

	if err != nil {
		panic("Can't auto migrate")
	}

	fmt.Println("Auto migrate ok!")
}

func setupRoutes(app *fiber.App) {
	api_routes := app.Group("/api/v1/")
	api.SetupApiRoutes(api_routes)

	book_routes := app.Group("")
	book.SetupBookRoutes(book_routes)
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
