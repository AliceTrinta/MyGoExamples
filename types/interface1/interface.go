package main

import "fmt"

type sport interface {
	turnOnTurbo()
}

type ferrari struct {
	model string
	turbo bool
}

func(f *ferrari) turnOnTurbo(){
	f.turbo = true
}

func main() {
	car1 := ferrari{"F40", false}
	car1.turnOnTurbo()
	var car2 sport = &ferrari{"F40", false}
	car2.turnOnTurbo()
	fmt.Println(car1, car2)
}
