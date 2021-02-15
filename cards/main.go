package main

import "fmt"

func main() {
	// var card string = "Ace of spades" # Full declaration
	card := newCard()

	fmt.Println(card)
	fmt.Println(printState())

	// Array -> fixed lenght
	// Slice -> An array that can growth
	cards := deck{"Five of Diadmonds", "Ace of Spades", newCard(), newCard()}
	cards = append(cards, "Six of Spades") // creates a new Array (Inmutability)

	fmt.Println(cards)
	cards.print()

	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}

type laptopSize float

func (this laptopSize) get() {
	return this
}
