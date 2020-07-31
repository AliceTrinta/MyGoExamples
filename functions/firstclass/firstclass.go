package main

import "fmt"

//passing anonymous function result to be stored in a variable.
//Another way to declare functions
var plus = func(a, b int) int{
	return a+b
}

func main(){
	fmt.Println(plus(5, 2))

	min := func(a, b int) int{
		return a-b
	}

	fmt.Println(min(5, 2))
}