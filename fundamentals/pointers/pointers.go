package main

import "fmt"

func main()  {
	i := 1
	var p *int = nil

	//Catch the reference.
	p = &i

	//Here we use "*" to catch the content in memory.
	*p++
	//The same as doing...
	i++

	fmt.Println("To stay clear:", i, *p, p, &p)

	//Go doesn't have pointers arithmetic.
	//E.g. ++p --> this is wrong.

}
