package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	lastName string
}

func (p person) getCompleteName() string{
	return p.name + " " + p.lastName
}

//Passing a pointer to a person struct
func (p *person) setCompleteName(CompleteName string){
	pieces := strings.Split(CompleteName, " ")
	p.name = pieces[0]
	p.lastName = pieces[1]
}

func main() {
	p1 := person{"Pietro", "Mendoza"}
	fmt.Println(p1.getCompleteName())
	p1.setCompleteName("James Dean")
	fmt.Println(p1.getCompleteName())
}
