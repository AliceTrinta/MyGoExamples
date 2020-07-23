package main

import "fmt"

func main(){
	s := make([]int, 10) //Slice from make function.
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) //Internal array now have 20 positions.
	fmt.Println(s, len(s), cap(s)) //Slice, slice length and internal array length or capacity.

	s = append(s, 1,2,3,4,5,6,7,8,9,0) //Adding elements to the slice, automatically to the internal array.
	fmt.Println(s, len(s), cap(s)) //Slice, slice length and internal array length or capacity.

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s)) //When the capacity is overcome, the internal array grow automatically.
}
