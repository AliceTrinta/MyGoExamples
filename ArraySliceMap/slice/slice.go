package main

import "fmt"

func main()  {
	a1 := [3]int{1,2,3} //Array.
	s1 := []int{1,2,3} //slice.

	fmt.Println(a1, s1)

	a2 := [5]int{1,2,3,4,5}
	//Slice is a piece of an array. Slice points, Array store.
	s2 := a2[1:3]
	s3 := a2[:2] //New slice, points the same array (a2).

	fmt.Println(a2, s2, s3)

	s4 := s2[:1] //Slice from a slice
	fmt.Println(s2, s4)


}
