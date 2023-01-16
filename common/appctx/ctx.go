package appctx

import "github.com/gofiber/fiber/v2"

type Context struct {
	*fiber.Ctx
}

func Wrap(h func(*Context) error) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return h(&Context{
			Ctx: ctx,
		})
	}
}
