package main

import "fmt"

//Creating a func that returns a func.
func closure() func(){
	x := 10
	//Creating a func under this context.
	//Example of Anonymous function.
	var function = func() {
		fmt.Println(x)
	}
	return function
}
//This is closure.

func main(){
	x := 20
	fmt.Println(x)
	printx := closure()
	printx()

	//It print's 10 because the func knows the
	//context where it was created, and respects that.
}