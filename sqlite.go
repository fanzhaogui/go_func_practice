package main

import (
	"database/sql"
	 _ "github.com/mattn/go-sqlite3"
	"log"
)

func main()  {
	db, err := sql.Open("sqlite3", "file:locked.sqlite?cache=shared")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(1)

	// db.Query("")
}
