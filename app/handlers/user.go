package handlers

import (
	"github.com/blog-master/pkg/appctx"
)

func Login(ctx *appctx.Context) appctx.ApiError {
	return ctx.ResponseBadRequest("deneme")
}

func Logout(ctx *appctx.Context) appctx.ApiError {
	return ctx.ResponseSuccess(nil)
}
