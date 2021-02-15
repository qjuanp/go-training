package main

import "fmt"

func main() {
	// var card string = "Ace of spades" # Full declaration
	card := "Ace of spades"   // Short declaration and assignation with type inferance
	card = "Five of Diamonds" // Assignation

	fmt.Println(card)
}
