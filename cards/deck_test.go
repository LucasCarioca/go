package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d:=newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck len of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spades"{
		t.Errorf("Expected first card to be Ace of Spades but got %s", d[0])
	}

	if d[len(d)-1] != "King of Clubs"{
		t.Errorf("Expected last card to be King of Clubs but got %s", d[51])
	}

	os.Remove("_decktest")
	d.saveToFile("_decktest")
	nd, _ := loadFromFile("_decktest")
	if len(nd) != 52 {
		t.Errorf("Expected deck from file to be of length 52 but got %v", len(nd))
	}
	os.Remove("_decktest")

 
}
