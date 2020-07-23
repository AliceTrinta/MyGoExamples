package main

import "fmt"

func main(){
	//var list map[int] string
	//A map have to be initialized.
	list := make(map[int]string)

	list[1]="Ana"
	list[1]="Lisa" //Only the last declaration will be valid.
	list[2]="Charles"
	list[3]="Jolene"
	fmt.Println(list)

	for id, name := range list{
		fmt.Printf("%s (ID: %d).\n", name, id)
	}

	fmt.Println(list[1])
	delete(list, 1) //If deleting, the content related to the key will be null.
	fmt.Println(list[1])

}
