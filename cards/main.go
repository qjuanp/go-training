package main

import "fmt"

func main() {
	// var card string = "Ace of spades" # Full declaration
	cards := newDeck()
	fmt.Println(cards.toString())
}
