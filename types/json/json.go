package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	//Putting the first letter in uppercase to show that
	//the variable is public.
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Tags []string `json:"tags"`
}

func main() {
	//Struct to json
	p1 := product{1, "Notebook", 10.2, []string{"Top10", "Most wanted"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	//json to struct
	var p2 product
	jsonString := `{"id":2,"name":"Pencil","price":1.9,"tags":["New", "Favorites"]}`
	//transforming jsonString into a slice of byte that can be stored in p2.
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Price)
}
