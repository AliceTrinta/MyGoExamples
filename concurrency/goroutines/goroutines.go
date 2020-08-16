package main

import (
	"fmt"
	"time"
)

func talk(person, text string, quantity int){
	for i := 0; i < quantity; i++{
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteraction %d)\n", person, text, i+1)
	}
}

func main(){
	//Serial way, once at a time.
	//talk("Jolene", "Why you're not speaking?", 3)
	//talk("Leon", "I just can talk after you!", 1)

	//Creating goroutine with the reserved word "go".
	//Executing the code in a independent way, like a thread but
	//lighter, with less resources.

	//go talk("Mary", "Hey!", 500)
	//go talk("Jorge", "Hello!", 500)

	//When executing the main, one line of execution is created,
	//from this line, multiple other lines are created.
	//In go we create goroutines, not threads! (Functions executed
	// in a independent way)
	//If the main ends first, the goroutine ends together. Differently when
	//the main waits the end of all the process under it self.
	//So, to make talk func work, the main func have to be more patient
	//time.Sleep(time.Second * 5)

	//go talk("Mary", "Ok!", 10)
	//talk("Jonas", "Good!", 5)
	//When doing this, the first talk will be executed independently
	//while the second will be in the line of execution of main func,
	//When the second func end, the main ends and it doesn't matter if the
	//first talk isn't finished yet.

	//Concluding, calling "go" to some func make it's execution independent.
}