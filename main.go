package main

import (
	"database/sql"
	"encoding/json"
	"log"

	"/Users/matt/dev/go-api/sqlusers"

	_ "github.com/mattn/go-sqlite3"
	"github.com/valyala/fasthttp"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(sqlusers.CreateUsersTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(sqlusers.InsertUsers)
	if err != nil {
		log.Fatal(err)
	}

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/users":
			getUsers(ctx, db)
		default:
			ctx.Error("Invalid path", fasthttp.StatusNotFound)
		}
	}

	fasthttp.ListenAndServe(":3000", requestHandler)
}

func getUsers(ctx *fasthttp.RequestCtx, db *sql.DB) {
	rows, err := db.Query(sqlusers.SelectUsers)
	if err != nil {
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
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
