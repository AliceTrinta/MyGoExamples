package main

import (
	"fmt"
	m "math" //You can use nicknames for imports.
)

func main(){
	const pi float64 = 3.1415
	var raid = 3.2 //float64 type was induced.

	//reduced form to creat a variable.
	//m for math.
	area := pi * m.Pow(raid, 2)

	//The program will compile only if you use all the variables.
	fmt.Println("Area is", area)

	//other forms to declare variables are...
	const(
		a = 1
		b = 2
	)
	const(
		c = 3
		d = 4
	)
	var e, f bool = true, false
	g, h, i := 2, true, "JustAVariable"

	//Just some prints
	fmt.Println("Numbers are is", a,b,c,d)
	fmt.Println("Booleans", e, f)
	fmt.Println("More variables", g,h,i)
}
