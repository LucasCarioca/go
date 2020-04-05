package main

import (
	"fmt"
)

func main() {

	//try to load hand and deck
	hand, err := loadFromFile("Hand.cards")
	if err != nil {
		fmt.Println(err)
	}
	cards, err := loadFromFile("Cards.cards")
	if err != nil {
		fmt.Println(err)
		//if hand and deck cant be loaded create a new deck and deal a hand
		hand, cards = newDeck().shuffle().deal(5)
	}

	fmt.Println("-- Hand --")
	hand.print()

	fmt.Println("-- Deck --")
	cards.print()
	// hand.saveToFile("Hand.cards")
	// cards.saveToFile("Cards.cards")
}
