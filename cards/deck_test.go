package main

import (
	"os"
	"testing"
)

func TestNewDEck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length 16 but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected card 'Ace of Spades' but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected card 'Four of Clubs' but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFilename := "_deckTestFile"
	os.Remove(testFilename)

	deck := newDeck()
	err := deck.saveToFile(testFilename)

	if err != nil {
		t.Error("There was an error saving the file", err)
	}

	deckFromFile := newDeckFromFile(testFilename)

	if len(deck) != len(deckFromFile) {
		t.Errorf("Expected deck size of %d but got %d", len(deck), len(deckFromFile))
	}

	if deck[0] != deckFromFile[0] {
		t.Errorf("Expected card %s but got %s", deck[0], deckFromFile[0])
	}

	if deck[len(deck)-1] != deckFromFile[len(deckFromFile)-1] {
		t.Errorf("Expected card %s but got %s", deck[len(deck)-1], deckFromFile[len(deckFromFile)-1])
	}

	os.Remove(testFilename)
}
