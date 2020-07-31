package main

import "fmt"

//Creating a func that doesn't have a number os parameters defined.
func avarage(numb...float64) float64{
	total := 0.0
	for _, num := range numb{
		total += num
	}
	return total/float64(len(numb))
}

func main()  {
	fmt.Printf("The avarage is: %.2f", avarage(1,2,3))
}
