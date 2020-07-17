package main

import (
	"fmt"
	"math"
	"reflect"
)

func main()  {
	//Integers
	fmt.Println(1, 2, 1000)
	fmt.Println("Integer's literal type is:", reflect.TypeOf(-32000))

	//Only positive integers: uint8(byte), uint16, uint32, uint64.
	var b byte = 255
	fmt.Println("Byte's type is:", reflect.TypeOf(b))

	//With signal. Can be negative: int8, int16, int32, int64.
	i1 := math.MinInt64
	fmt.Println("The min value is:", i1)
	fmt.Println("The type of i1 is:", reflect.TypeOf(i1))

	//Runes represents the respective value of a variable on the unicode's table.
	var i2 rune = 'a'
	fmt.Println("Rune's type is:", reflect.TypeOf(i2))
	fmt.Println("Rune's value is", i2)

	//Real numbers: float32, float64.
	var x float32 = 32.332
	fmt.Println("var float type is:", reflect.TypeOf(x))
	fmt.Println("float literal type is:", reflect.TypeOf(32.332))

	//Boolean
	bo:=true
	fmt.Println("Boolean's are", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//String
	s1 := "Only a Sting"
	fmt.Println(s1 + "!")
	fmt.Println("lenght:", len(s1))

	//String with multiple lines.
	s2 := `Backtick
		with
		multiple
		lines`
	fmt.Println("lenght:", len(s2))

	//Char (?)
	//var letter char = 'a'
	letter := 'a'
	fmt.Println("Char type is:", reflect.TypeOf(letter))
	//With char we have a int32 type that's associated with the unicode's table.

}
