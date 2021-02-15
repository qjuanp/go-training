package main

import "fmt"

func main() {
	// var card string = "Ace of spades" # Full declaration
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	fmt.Println("--Hand")
	hand.print()
	fmt.Println("--remainingCards")
	remainingCards.print()
}
