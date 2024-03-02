package main

import (
	"os"
	"testing"
)

const DeckSize = 52

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != DeckSize {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades'")
	}

	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected last card to be 'King of Clubs'")
	}
}

func TestSavetoDeckAndNewDeckTestFromFile (t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != DeckSize {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	os.Remove("_decktesting")
}

