package main

import (
	"fmt"
	"strconv"
)

func main()  {
	x := 2.4
	y := 2

	//When doing operations, you need to use the same data type.
	fmt.Println(x/float64(y))
	//fmt.Println(int(x)/y)

	points := 6.9
	finalPoints := int(points)
	fmt.Println(finalPoints)

	//Caution
	fmt.Println("Final points:" + string(123)) //Unicode table for code 123!

	//Integers to string
	fmt.Println("Final points:" + strconv.Itoa(123))

	//String to integer
	num, _ := strconv.Atoi("123") //This function returns an int and an error. To ignore a value, type "_".
	fmt.Println("Final points:", num-122)

	//String to bool
	b, _ := strconv.ParseBool("1")
	if b{
		fmt.Println("True")
	}

}
