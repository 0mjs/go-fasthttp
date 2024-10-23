package seed

import (
	"database/sql"
	"log"

	usersSQL "go-api/sql"
)

func SeedDB(db *sql.DB) {
	_, err := db.Exec(usersSQL.DropUsersTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(usersSQL.CreateUsersTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(usersSQL.InsertUsers)
	if err != nil {
		log.Fatal(err)
	}
}
