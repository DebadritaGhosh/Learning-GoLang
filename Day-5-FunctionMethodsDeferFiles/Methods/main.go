package main

import "fmt"

func main() {
	fmt.Println("Methods...")
	//No inheritance in golang
	//No super or parent

	debadrita := User{"Debadrita", "debadrita.ghosh2010@gmail.com", true, 12}
	fmt.Println(debadrita)
	fmt.Printf("Debadrita's details are : %+v\n", debadrita)
	fmt.Printf("Debadrita's email is %v and status is %v \n", debadrita.Email, debadrita.Status)
	debadrita.GetStatus()
	debadrita.newMail()
	fmt.Println(debadrita.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) newMail() {
	u.Email = "debadrita@gmail.com"
	fmt.Println("Email of this user is : ", u.Email)
}
