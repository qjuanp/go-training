package main

import "fmt"

func main() {
	// var card string = "Ace of spades" # Full declaration
	card := newCard()

	fmt.Println(card)
	fmt.Println(printState())
}

func newCard() string {
	return "Five of Diamonds"
}
