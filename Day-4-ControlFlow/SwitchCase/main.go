package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch & Case in GoLang...")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 step")
	case 3:
		fmt.Println("You can move 3 step")
		// fallthrough
	case 4:
		fmt.Println("You can move 4 step")
	case 5:
		fmt.Println("You can move 5 step")
	case 6:
		fmt.Println("You can move 6 step and Roll the dice again")
	default:
		fmt.Println("What was that!")
	}

}
