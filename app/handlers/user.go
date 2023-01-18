package handlers

import (
	"github.com/blog-master/pkg/appctx"
)

func Login(ctx *appctx.Context) appctx.ApiError {
	return ctx.ResponseBadRequest("xd")

}

func Logout(ctx *appctx.Context) appctx.ApiError {
	return ctx.ResponseSuccess("çıkış yaptınız")
}
