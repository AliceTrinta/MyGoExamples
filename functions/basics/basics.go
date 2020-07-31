package main

import "fmt"

func f1()  {
	fmt.Println("F1")
}

func f2(p1 string, p2 string)  {
	fmt.Println("F2:", p1, p2)
}

func f3() string  {
	return "F3"
}

func f4(p1, p2 string) string{
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

//It's possible have functions with many parameters.
func f5() (string, string) {
	return "1", "2"
}

func main() {
	f1()
	f2("ola", "hi")

	a3, a4 := f3(), f4("ciao", "hallo")
	fmt.Println(a3)
	fmt.Println(a4)

	//It's possible ignore one of the returns.
	//a51, a52 := f5()
	_, a52 := f5()
	fmt.Println(a52)
}
