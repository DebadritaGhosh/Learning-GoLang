// With extra variable
package main

import "fmt"

func main() {
	var numOne int = 23
	var numTwo int = 26
	var temp int
	temp = numOne
	numOne = numTwo
	numTwo = temp
	fmt.Printf("Number One is %d number Two is %d", numOne, numTwo)
}

// No extra variable
func main() {
	var numberOne int = 36
	var numberTwo int = 24
	numberOne = numberOne + numberTwo
	numberTwo = numberOne - numberTwo
	numberOne = numberOne - numberTwo
	fmt.Printf("Number One is %d & number Two is %d", numberOne, numberTwo)
}
