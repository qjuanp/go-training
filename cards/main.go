package main

import "fmt"

func main() {
	// var card string = "Ace of spades" # Full declaration
	card := newCard()

	fmt.Println(card)
	fmt.Println(printState())

	// Array -> fixed lenght
	// Slice -> An array that can growth
	cards := []string{"Five of Diadmonds", "Ace of Spades", newCard(), newCard()}
	cards = append(cards, "Six of Spades") // creates a new Array (Inmutability)

	fmt.Println(cards)

	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
