package main

import "fmt"

func printAccepted(list...string) {
	fmt.Println("List of accepted")
	for i, name := range list{
		fmt.Printf("(%d) %s\n", i+1, name)
	}
}

func main(){
	//Creating a slice with the content.
	accepted := []string{"Mary", "Rob", "Bob"}
	//Passing the slice as a parameter.
	printAccepted(accepted...)
	//Only work with slices, when nether the func and the slice know
	//the number of contents.
}