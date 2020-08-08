package main

import "fmt"

//Here we create a custom type that will have it's own methods
type result float64

func (r result) under(begin, end float64) bool{
	return float64(r) >= begin && float64(r) <= end

}

func conclusion(r result) string{
	if r.under(9.0, 10.0){
		return "A"
	} else if r.under(7.0, 8.99){
		return "B"
	} else if r.under(5.0, 7.99) {
		return "C"
	} else if r.under(3.0, 4.99){
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(conclusion(6.7))
}
