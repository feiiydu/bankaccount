package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	const connStr = "postgres://uzzyzabm:6JOV9j9Ov1utIojEfRsRgU57wMz9r7X6@elmer.db.elephantsql.com:5432/uzzyzabm"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
