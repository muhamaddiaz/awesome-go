package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 9 {
		t.Errorf("Expected 9 values, but got %v", len(cards))
	}

	if cards[0] != "Hi, Diaz" {
		t.Errorf("Expected Hi, Diaz, but got %v", cards[0])
	}

	if cards[len(cards) - 1] != "Wow, Abigail" {
		t.Errorf("Expected Wow, Abigail, but got %v", cards[len(cards) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_testcards")

	cards := newDeck()
	cards.writeDeckToFile("_testcards")

	if len(readDeckFromFile("_testcards")) != 9 {
		t.Errorf("Reading files expect 9 values, but got %v", len(readDeckFromFile("_testcards")))
	}

	os.Remove("_testcards")
}
