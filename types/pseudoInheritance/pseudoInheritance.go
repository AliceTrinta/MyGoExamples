package main

import "fmt"

type car struct {
	name string
	price float64
}

/*Declaring a struct under a struct without initialize
variables will create anonymous fields. So, everything
that a car has, a ferrari can have too.*/
type ferrari struct {
	car //Just a composition, seems like inheritance but is not.
	turbo bool
}

func main() {
	f := ferrari{}
	f.name = "F40" //From car
	f.price = 100000000.0 //From car
	f.turbo = true

	fmt.Printf("Ferrari %s that costs %.2f has turbo? %v", f.name, f.price, f.turbo)
	fmt.Println(f) //This shows that we only have a composition, not inheritance.
}
