package main

import "fmt"

func main() {
	funcsByLetter := map[string]map[string]float64{
		"G" : {
			"Giana" : 1200.23,
			"Gustav" : 1330.32,
		},
		"B" : {
			"Bela" : 1200.23,
			"Boston" : 1330.32,
		},
	}

	for letter, funcs := range funcsByLetter{
		fmt.Println(letter, funcs)
	}


}
