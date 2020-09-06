package math //To run all tests under package.

import "testing"

const error = "Unexpected value %v. Expected value was %v."

//The name in thus test function is following the patterns given by the language.
func TestAvarage(t *testing.T) { //To run only this test.
	expectedValue := 7.28
	value := Avarage(7.2, 9.9, 6.1, 5.9)
	if value != expectedValue{
		t.Errorf(error, value, expectedValue)
	}
}
