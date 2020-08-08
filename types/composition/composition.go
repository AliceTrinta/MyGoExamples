package main

import "fmt"

type sport interface {
	turnOnTurbo()
}

type lux interface {
	automaticParking()
}

//Interface that's composed of interfaces
type sportDelux interface {
	sport
	lux
	//Can add more
}

type bmw struct {}

func (b bmw) turnOnTurbo(){
	fmt.Println("Turning on turbo mode...")
}

func (b bmw) automaticParking(){
	fmt.Println("Parking...")
}

func main() {
	var b sportDelux = bmw{}
	b.turnOnTurbo()
	b.automaticParking()
}
