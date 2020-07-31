package main

import "fmt"

func value(n int) int{
	defer fmt.Println("END")
	if n >= 5000 {
		fmt.Println("Good value")
		return 5000
	}
	fmt.Println("Bad value")
	return n
}

func main(){
	fmt.Println(value(400))
}