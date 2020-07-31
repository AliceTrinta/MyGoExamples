package main

import "fmt"

func exchange(p1, p2 int) (second int, first int){
	second = p2
	first = p1
	return //Clean return
}

func main() {
	a1, a2 := exchange(1, 2)
	fmt.Println(a1, a2)
}
