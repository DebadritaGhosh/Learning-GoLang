package main

import "fmt"

func main() {
	fmt.Println(" Welcome to pointers! ")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 26
	var ptr = &myNumber
	fmt.Println("Value of ptr is : ", ptr)
	fmt.Println("Value of *ptr is : ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is : ", myNumber)

}
