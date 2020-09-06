package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
		db, err := sql.Open("mysql", "root:Alice090@/goworkshop")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into users(id, name) values(?, ?)")
	stmt.Exec(2000, "Charles")
	stmt.Exec(2001, "Leon")
	_, err = stmt.Exec(4, "Melanie") //Duplicated key

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
}
