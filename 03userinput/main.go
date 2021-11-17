package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter Your Age")

	// Comma Ok Syntax || Error Ok Syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Your Age Is : ", input)
	fmt.Printf("Type of this rating is %T", input)
}
