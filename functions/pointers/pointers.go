package main

import "fmt"

//Only create a copy of the variable.
func add1(n int){
	n++
}

//Receiving a pointer to an integer.
func add2(n *int){
	//* is used here to get the content and adding one to the variable.
	*n++
}

func main(){
	n:=1
	add1(n)
	fmt.Println(n)

	//& is used to get the path of the variable
	add2(&n)
	fmt.Println(n)

	//Be careful! Avoid pass the reference, try to pass only values.
}