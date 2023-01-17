package appctx

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type Context struct {
	*fiber.Ctx
	//ApiError
}

func Wrap(h func(*Context) ApiError) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return h(&Context{
			Ctx: c,
		})
	}
	//todo: içeride sıfır bir fiber ctx dönüp kendi ctx'imde o fiber ctx'ini mutate etmiyor mıyum ben ya. neden status code'lerini falan almıyor?  alıyor aslında ama postman'da hata kodu 500 gösteriyor ne olursa olsun. belki de bu Error methodu ile alakalı olabilir. bilmiyorum.
}

type ApiError struct {
	Data interface{} `json:"data"`
	//Message string      `json:"message"`
}

// todo: çalışıyor ama yanlış geliyor nedense.
func (c ApiError) Error() string {
	a, _ := json.Marshal(c)
	return string(a)
}

func (c Context) ResponseSuccess(data interface{}) ApiError {
	c.Status(200)
	return ApiError{
		Data: data,
	}
}

func (c Context) ResponseBadRequest(data interface{}) ApiError {
	c.Status(400)
	return ApiError{
		Data: data,
	}
}
