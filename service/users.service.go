package service

import (
	"database/sql"
	"encoding/json"

	"go-api/model"
	usersSQL "go-api/sql"

	"github.com/valyala/fasthttp"
)

func FetchUsers(ctx *fasthttp.RequestCtx, db *sql.DB) {
	rows, err := db.Query(usersSQL.SelectUsers)
	if err != nil {
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	response, err := json.Marshal(users)
	if err != nil {
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(response)
}
