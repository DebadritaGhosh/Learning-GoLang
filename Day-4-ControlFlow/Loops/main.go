package main

import "fmt"

func main() {
	fmt.Println("Loops............")

	days := []string{"Sunday", "Tuesday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println((days[i]))
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v \n", index, day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto hello
		}

		if rougueValue == 5 {
			break
		}

		fmt.Println("Value is : ", rougueValue)
		rougueValue++
	}

	//goto
hello:
	fmt.Println("Hello World")

}
