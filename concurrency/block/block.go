package main

import (
	"fmt"
	"time"
)

func routine(c chan int){
	time.Sleep(time.Second)
	c <- 1 //Block operation, waiting someone read.
	fmt.Println("Only after the channel!")
}

func main() {
	c := make(chan int) //Channel without buffer
	//When the channel doesn't have buffer, you can only pass
	//a value when the first was read.

	go routine(c)

	fmt.Println(<-c) //Block operation, will wait the data be ready
	fmt.Println("Was read!")
	fmt.Println(<-c) //deadlock
	fmt.Println("This is not going to be appear because of deadlock")
}
