package main

import (
	"github.com/jmoiron/sqlx"
)

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

func main() {
	db, _ := sqlx.Connect("")

	people := []Person{}
	db.Select(&people)
}
