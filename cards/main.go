package main

import "fmt"

// Gets executed when running `go run main.go`
func main() {
	// declares and assigns a value to variable card
	// var card string = "Ace of Spades"
	// := makes Go, figure out what type of variable we are declaring
	card := createCard()
	// we can also reassign this variable to another string by --
	// card = "Any string, notice dropping the :"
	fmt.Println(card)
}

// requires you tell it what is should return
func createCard() string {
	return "Five of Diamonds"
}
