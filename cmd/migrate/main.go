package main

import "github.com/monci/internal/db"

func main() {
	db := db.Db{}
	err := db.Open()
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}
}
