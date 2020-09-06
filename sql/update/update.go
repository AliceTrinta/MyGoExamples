package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" )

func main() {
	db, err := sql.Open("mysql", "root:Alice090@/goworkshop")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Update
	stmt, _ := db.Prepare("update users set name = ? where id = ?")

	stmt.Exec("Messi", 1)
	stmt.Exec("Neymar", 2)

	//Delete
	stmt2, _ := db.Prepare("delete from users where id = ?")
	stmt2.Exec(3)
}