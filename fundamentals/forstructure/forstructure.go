package main

import (
	"fmt"
	"time"
)

func main()  {
	//For is the only repetition structure in Go!
	//Declaring classic "for".
	for i := 1;i < 10; i++{
		fmt.Println("First for", i)
	}

	//Declaring "for" in "while" style.
	i := 0;
	for i < 10{
		fmt.Println("Second for", i)
		i++
	}

	//Declaring an infinity "for".
	for{
		fmt.Println("Infinity for")
		time.Sleep(time.Second * 5)
	}
}
