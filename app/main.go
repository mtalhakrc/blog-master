package main

import (
	"github.com/blog-master/app/router"
	"github.com/blog-master/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	router.Setup(app)

	app.Listen(":8080")
}
