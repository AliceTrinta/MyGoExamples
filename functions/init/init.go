package main

import "fmt"

//It's always initialized before the main of the program.
func init()  {
	fmt.Println("Executing...")
}

func main()  {
	fmt.Println("Here is main!")
}