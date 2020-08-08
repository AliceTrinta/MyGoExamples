package main

import "fmt"

//If we create a structure that allows interface implementation
//The compiler will recognize that turning easier the use of
//polymorphism. Interfaces are implemented implicitly.

type printable interface {
	toString() string
}

type person struct {
	name string
	lastName string
}

type product struct {
	name string
	price float64
}

func (p person) toString() string{
	return p.name + " " + p.lastName
}

func (p product) toString() string{
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(x printable){
	fmt.Println(x.toString())
}

//If the methods of a struct matches with the methods of an interface
//The struct will implement automatically that interface

func main() {
	var thing printable = person{"James", "Cornell"}
	fmt.Println(thing.toString())
	print(thing)
	//Polymorphism based on interfaces.
	thing = product{"Pencil", 1.2}
	print(thing)
	//If respects the structure of the interface, can be an interface.
	thing2 := product{"Tomato", 2.0}
	print(thing2)
}
