package main

import "fmt"

func main()  {
	//Channel is a type, like int strings etc. To create use make func and declare
	//the type of data that is going to be passed.
	ch := make(chan int, 1)

	ch <- 1 //Sending data to channel (writing)
	<- ch //Receiving data (reading)

	ch <- 2
	fmt.Println(<-ch)

	//A channel is a point of synchronism
}
