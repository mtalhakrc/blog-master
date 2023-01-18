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
		//return errors.New("demekki handler'lar error interfacesinin Error()string metodununu çalıştırarak cevap olrak string dönyor. ama bu şekilde error dönünce otomatik olarak internal server code'sini yazıyor. o halde response body'e yazmak lazım error olarak nil return etmeli ve status codelerini manuel olarak vermeli.")
		return nil
	})
}
