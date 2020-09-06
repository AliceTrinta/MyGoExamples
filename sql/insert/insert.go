package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {

	db, err := sql.Open("mysql", "root:Alice090@/goworkshop")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, _ := db.Prepare("insert into users(name) values(?)")
	stmt.Exec("Mary")
	stmt.Exec("Jonesy")

	ans, _ := stmt.Exec("Mary")
	id, _ := ans.LastInsertId()
	fmt.Println(id)

	lines, err := ans.RowsAffected()
	fmt.Println(lines)
}
