package main

import "fmt"

func main(){
	var result [3] float64
	fmt.Println(result)

	result[0], result[1], result[2] = 7.0, 1.2, 5.0
	fmt.Println(result)

	total:=0.0

	for i:=0; i<len(result); i++{
		total=+result[i]
	}

	avarage := total/float64(len(result))

	fmt.Printf("%2.f", avarage)

}