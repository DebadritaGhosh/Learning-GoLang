package main

import "fmt"

func main() {
	fmt.Println("Functions...........")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is :", result)

	proResult, myMessage, myMessage2 := proAdder(21, 6, 5, 4, 7, 3, 5)
	fmt.Println("Result is :", proResult)
	fmt.Println("Message is :", myMessage)
	fmt.Println("Message is :", myMessage2)

}

func greeter() {
	fmt.Println("Namstey from golang!")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hiiii", "Hello"
}
