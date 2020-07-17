package main

import (
	"fmt"
	"time"
)

func main()  {

	//Demonstration of the switch structure in go.
	t := time.Now()

	switch {
	case t.Hour() > 10 && t.Hour() <17:
		i := 10.0
		fmt.Println("You get:", getResults(i))
	default:
	}
}

func getResults(result float64) string {
	//Switch in Go is different than in other languages...
	switch result {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8,7:
		return "B"
	case 6,5:
		return "C"
	case 4,3:
		return "D"
	case 2,1,0:
		return "E"
	default:
		return "Invalid Result!"
	}
}

