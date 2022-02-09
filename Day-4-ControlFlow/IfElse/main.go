package main

import "fmt"

func main() {
	fmt.Println("IF Else in GoLang..")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10"
	}
	fmt.Println(result)

	fmt.Println("---------------------------")

	//Assigning and checking the value at the same time
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}
}
