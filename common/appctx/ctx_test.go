package appctx

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http/httptest"
	"testing"
)

func okHandler(ctx *Context) error {
	return ctx.SendString("ok")
}
func notokHandler(ctx *Context) error {
	ctx.Status(400)
	return ctx.SendString("not ok")
}

type isimbulamadım struct {
	method           string
	route            string
	expectedResponse string
	expectedStatus   string
	reqBody          io.Reader

	handler func(ctx *Context) error
}

func TestCtx(t *testing.T) {
	scenerios := []isimbulamadım{
		{
			method:           "GET",
			route:            "/get",
			expectedResponse: "ok",
			expectedStatus:   "200 OK",
			reqBody:          nil,
			handler:          okHandler,
		},
		{
			method:           "POST",
			route:            "/post",
			expectedResponse: "ok",
			expectedStatus:   "200 OK",
			reqBody:          &bytes.Buffer{},
			handler:          okHandler,
		},
		{
			method:           "PUT",
			route:            "/put",
			expectedResponse: "not ok",
			expectedStatus:   "400 Bad Request",
			reqBody:          nil,
			handler:          notokHandler,
		},
		{
			method:           "DELETE",
			route:            "/delete",
			expectedResponse: "not ok",
			expectedStatus:   "400 Bad Request",
			reqBody:          nil,
			handler:          notokHandler,
		},
	}

	app := fiber.New()

	for _, scene := range scenerios {
		switch scene.method {
		case "GET":
			app.Get(scene.route, Wrap(scene.handler))
		case "POST":
			app.Post(scene.route, Wrap(scene.handler))
		case "PUT":
			app.Put(scene.route, Wrap(scene.handler))
		case "DELETE":
			app.Delete(scene.route, Wrap(scene.handler))
		}

		req := httptest.NewRequest(scene.method, scene.route, scene.reqBody)
		res, _ := app.Test(req, 1)
		str, _ := io.ReadAll(res.Body)
		if string(str) != scene.expectedResponse {
			t.Errorf("got: %s; want: %s", string(str), scene.expectedResponse)
		}
		if res.Status != scene.expectedStatus {
			t.Errorf("got: %s; want: %s", res.Status, scene.expectedResponse)
		}
	}

}
