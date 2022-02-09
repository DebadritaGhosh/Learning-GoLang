package main

import "fmt"

func main() {
	fmt.Print("MAPS")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	//Printing all maps
	fmt.Println("List of all languages : ", languages)
	//Printing single value
	fmt.Println("JS stands for : ", languages["JS"])
	//Delete
	delete(languages, "RB")
	fmt.Println(languages)
}
