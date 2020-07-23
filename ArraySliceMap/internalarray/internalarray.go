package main

import "fmt"

func main(){
	//To prove that slices can point to the same array, and change this array can change both slices.
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	s1[0] = 7
	fmt.Println(s1, s2)

}
