package main

import (
	"fmt"
)

func factorial(n uint) uint{
	switch {
	case n == 0:
		return 1
	default:
		return n * factorial(n-1)
	}
}

func main(){
	result := factorial(20)
	fmt.Println(result)

	//result := factorial(-20) -> It's not possible. The func only
	//receive uint values, integers without sign.
}
