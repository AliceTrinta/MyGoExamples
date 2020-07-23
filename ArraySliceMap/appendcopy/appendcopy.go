package main

import "fmt"

func main(){
	a1 := [3]int {1,2,3} //When declaring length, you know is an array.
	var s1 []int

	s1 = append(s1, 1, 2, 3) //Changes length, assumes that you're adding something.
	fmt.Println(a1, s1)

	s2 := make([]int, 2)
	copy(s2, s1) //Doesn't change length
	fmt.Println(s2)
}
