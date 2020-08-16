package main

import (
	"fmt"
	"time"
)

//Channel - Way to goroutines communicate
//A channel is a type.

func multi(base int, c chan int){
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second * 3)
	c <- 4 * base
}

func main(){
	c := make(chan int)
	go multi(2, c)

	a, b := <-c, <-c //Receiving data from channel
	fmt.Println(a, b)
	fmt.Println(<-c)

	//If you try to receive a data that doesn't exist
	//you will get a deadlock error.
}
