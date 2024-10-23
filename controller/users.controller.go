package controller

import (
	"database/sql"

	service "go-api/service"

	"github.com/valyala/fasthttp"
)

func Controller(ctx *fasthttp.RequestCtx, db *sql.DB) {
	switch string(ctx.Path()) {
	case "/users":
		service.FetchUsers(ctx, db)
	default:
		ctx.Error("Invalid path", fasthttp.StatusNotFound)
	}
}
