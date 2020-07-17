package main

import "fmt"

func main(){

	//Here we have the basic types in the default mode.
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	//As you can see, they'r all set to null/empty/0 values.
	fmt.Printf("%v, %v, %v, %q, %v", a, b, c, d, e)
}
