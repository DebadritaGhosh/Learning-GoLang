package main

import "fmt"

func main() {
	fmt.Println("Structs...")
	//No inheritance in golang
	//No super or parent

	debadrita := User{"Debadrita", "debadrita.ghosh2010@gmail.com", true, 12}
	fmt.Println(debadrita)
	fmt.Printf("Debadrita's details are : %+v\n", debadrita)
	fmt.Printf("Debadrita's email is %v and status is %v \n", debadrita.Email, debadrita.Status)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
