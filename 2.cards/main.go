package main

import "fmt"

func main() {
	cards := newDeck()
	cards.toString()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	fmt.Println("--------")
	remainingCards.print()
	fmt.Println("--------")
	fmt.Println(cards.toString())

	cards.saveToFile("my_cards")

	newCardsFromFile := newDeckFromFile("my_cards")
	fmt.Println(newCardsFromFile)
}
