package appctx

import (
	"github.com/gofiber/fiber/v2"
)

type Context struct {
	*fiber.Ctx
	//ApiError
}

func Wrap(h func(*Context) ApiError) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return h(&Context{c})
	}
	//todo: içeride sıfır bir fiber ctx dönüp kendi ctx'imde o fiber ctx'ini mutate etmiyor mıyum ben ya. neden status code'lerini falan almıyor?  alıyor aslında ama postman'da hata kodu 500 gösteriyor ne olursa olsun. belki de bu Error methodu ile alakalı olabilir. bilmiyorum.
}

//type ApiError struct {
//	Data interface{} `json:"data"`
//	Message string      `json:"message"`
//}

type ApiError interface {
	Error() string
}

type ApiResponse struct {
	Data interface{} `json:"data"`
}

func (c ApiResponse) Error() string {
	return ""
}

func (c *Context) ResponseSuccess(data interface{}) ApiError {
	c.Status(200)
	m := ApiResponse{Data: data}
	err := c.JSON(m)
	if err != nil {
		panic(err)
	}
	return nil
}

func (c *Context) ResponseBadRequest(data interface{}) ApiError {
	c.Status(400)
	m := ApiResponse{Data: data}
	err := c.JSON(m)
	if err != nil {
		panic(err)
	}
	return nil

}
