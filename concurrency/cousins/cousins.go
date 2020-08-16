package main

import (
	"fmt"
	"time"
)

func isCousinPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0{
			return false
		}
	}
	return true
}

func cousinPrime(n int, c chan int){
	begin := 2
	for i := 0; i < n; i++{
		for cousin := begin; ; cousin++{
			if isCousinPrime(cousin) {
				c <- cousin
				begin = cousin + 1
				time.Sleep(time.Microsecond * 100)
				break
			}
		}
	}
	//Important to do not generate deadlocks.
	close(c)
}

func main() {
	c := make(chan int, 30)
	go cousinPrime(cap(c), c)
	for cousin := range c{
		fmt.Printf("%d ", cousin)
	}
	fmt.Println("The End!")
}
