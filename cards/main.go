package main

import "fmt"

func main() {
	// var card string = "Ace of spades" # Full declaration
	cards := newDeck()
	cards.saveToFile("my_cards")

	newCards := newDeckFromFile("my_cards")
	fmt.Println("---- New Deck from file")
	newCards.print()
}
