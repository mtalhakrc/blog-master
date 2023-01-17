package main

import (
	"github.com/blog-master/app/router"
	"github.com/blog-master/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(logger.New())
	router.Setup(app)

	app.Listen(":8080")
}
