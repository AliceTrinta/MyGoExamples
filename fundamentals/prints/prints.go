package main

import "fmt"

func main()  {

	//There's many ways to print something.
	fmt.Print("Same")
	fmt.Print("line\n")
	fmt.Println("Unique line.")

	x := 3.123234
	//This is not possible:
	//fmt.Println("The number is" + x)

	//So, we use...
	xs := fmt.Sprint(x)
	fmt.Println("The number is " + xs)
	//or...
	fmt.Println("The number is", x, "and we can also continue like this")

	//One interesting way is...
	fmt.Printf("The number is %f\n", x)
	//So, we can manipulate as we see...
	fmt.Printf("The number is %.2f", x)

	//With different types
	a:=1
	b:=2.123
	c:=true
	d:="Thats all"
	fmt.Printf("\n%d, %f, %.1f, %t, %s", a, b, b, c, d)
	//In a generic way, use "v":
	fmt.Printf("\n%.v, %v, %v, %v", a, b, c, d)
}
