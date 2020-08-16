package main

import "fmt"

func routine(c chan int){
	c <- 1
	c <- 2
	c <- 3
	fmt.Println("Executing!")
	c <- 4
	c <- 5
	c <- 6
}

func main() {
	//Channel with buffer = 3, only will accept
	//3 data until some data be consumed.
	c := make(chan int, 3)
	go routine(c)

	fmt.Println(<-c)
}
