package main

import "fmt"

type subject struct {
	name string
}

func main() {
	//Generic (empty) interfaces can receive any kind of types.
	//Seems like object from Java.
	var thing interface{} //Interface used as a type.
	fmt.Println(thing)

	thing = 3
	fmt.Println(thing)

	type dynamic interface{}
	//Creating a type of interface and using as a type for other variables
	var thing2 dynamic = "Ops!"
	fmt.Println(thing2)

	thing2 = true
	fmt.Println(thing2)

	thing2 = subject{"Math"}
	fmt.Println(thing2)
}
