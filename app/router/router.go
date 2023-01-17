package router

import (
	"github.com/blog-master/app/handlers"
	"github.com/blog-master/pkg/appctx"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/login", appctx.Wrap(handlers.Login))
	api.Get("/logout", appctx.Wrap(handlers.Logout))

	api.Get("/deneme", func(c *fiber.Ctx) error {
		c.Status(200)
		return nil
	})
}
