package main

import "fmt"

func main(){
	age := 17
	name := "Alice"

	fmt.Printf("My age is %v and my name is Alice!\n", age)
	ages := fmt.Sprint(age)
	fmt.Print("My age is " + ages + " and my name is Alice!\n")
	fmt.Println("My age is", age, "and my name is", name, "!")

}