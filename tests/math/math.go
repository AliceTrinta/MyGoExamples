package math

import (
	"fmt"
	"strconv"
)

func Avarage(numbers... float64) float64{
	total := 0.0
	for _, num := range numbers{
		total += num
	}
	final1 := total / float64(len(numbers))
	final2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", final1), 64)
	return final2
}
