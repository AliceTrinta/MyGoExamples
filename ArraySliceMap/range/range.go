package main

import "fmt"

func main()  {
	numbers:=[...]int{1,2,3,4,5} // counts automatically

	for i, number := range numbers{
		fmt.Println(i, number)
	}

	for _, number := range numbers{
		fmt.Println(number)
	}

}
