package main

import "fmt"

func main() {
	fmt.Println("Array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Orange"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	fmt.Println("----------------")

	var vegList = [3]string{"Potato", "Beans", "Carrot"}

	fmt.Println(vegList)
	fmt.Println(len(vegList))
}
