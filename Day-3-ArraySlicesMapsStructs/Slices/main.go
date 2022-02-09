package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")
	var fruitList = []string{"Apple", "Orange", "Grapes"}
	fmt.Printf("Types of fruitlist is %T", fruitList)
	fmt.Println()

	//adding values
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:]) // will print all fruits except the 1st ([Orange Grapes Mango Banana])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // will print the 2nd and 3rd ([Grapes Mango])
	fmt.Println(fruitList)

	// ---------------------------------------

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 653
	highScores[2] = 239
	highScores[3] = 963

	highScores = append(highScores, 555, 666, 777)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted : ", highScores) // will print sorted values

	//Remove value from slice based on index
	fmt.Println("-----------------------------")

	var courses = []string{"Reactjs", "JavaScript", "ReactNative", "Python", "Go"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
