package main

import "fmt"

func multi(a, b int) int{
	return a*b
}

//Passing a func as a parameter.
func exe(function func(int, int) int, a, b int) int{
	return function(a, b)
}

func main(){
	//Creating a variable to store the return of
	//the exe func, passing multi as parameter
	result := exe(multi, 3, 2)
	fmt.Println(result)
}
