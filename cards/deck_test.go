package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected card Ace of Spades but got : %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected card four of clubs but got : %v", d[len(d)-1])
	}
}
func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.saveToFile("_deckTesting")

	ld := newDeckFromFile("_deckTesting")

	if len(ld) != 16 {
		t.Errorf("Expected 16 but got %v", len(ld))
	}
	os.Remove("_deckTesting")
}
