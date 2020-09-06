package main

import (
	"database/sql"
	json2 "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//User type
type User struct {
	ID 	 int    `json:"id"`
	Name string `json:"name"`
}

//UserHandler, func to analise request and send to the correct func.
func UserHandler(w http.ResponseWriter, r *http.Request){
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method =="GET" && id > 0:
		userById(w, r, id)
	case r.Method == "GET":
		allUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Sorry...")
	}
}

func userById(w http.ResponseWriter, r *http.Request, id int){
	db, err := sql.Open("mysql", "root:Alice090@/goworkshop")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u User
	db.QueryRow("select id, name  from users where id = ?", id).Scan(&u.ID, &u.Name)

	json, _ := json2.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func allUsers(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql", "root:Alice090@/goworkshop")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, name from users")
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	json, _ := json2.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
