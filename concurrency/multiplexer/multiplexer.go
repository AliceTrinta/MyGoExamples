package main

import (
	"MyGoExamples/concurrency/html"
	"fmt"
)

//Concurrency pattern "Multiplexer"

func route(origin <-chan string, destiny chan string){
	for{
		destiny <- <-origin
	}
}

//Multiplexer function: Get together messages from different channels in only one.
func add(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go route(input1, c)
	go route(input2, c)
	return c
}

func main() {
	c := add(
		html.Title("https://www.google.com", "https://www.youtube.com"),
		html.Title("https://www.google.com.br/drive", "https://www.amazon.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
