package main

import (
	"database/sql"
	"log"

	"go-api/controller"
	seed "go-api/db"

	_ "github.com/mattn/go-sqlite3"
	"github.com/valyala/fasthttp"
)

func main() {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	seed.SeedDB(db)

	fasthttp.ListenAndServe(":3000", func(ctx *fasthttp.RequestCtx) {
		controller.Controller(ctx, db)
	})
}
