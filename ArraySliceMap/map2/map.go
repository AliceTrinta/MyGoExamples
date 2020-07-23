package main

func main() {
	funcsAndSalaries := map[string] int{
		"Ana" : 11200,
		"Lisa" : 12000,
		"James" : 13000,
	}

	funcsAndSalaries["Jonesy"] = 14000
	delete(funcsAndSalaries, "No one")

	for name, salary := range funcsAndSalaries{
		println(name, salary)
	}

}
