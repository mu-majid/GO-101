package main

import "fmt"

func main() {
	var card string = "Ace of Spades"
	anotherCard := newCard()
	// Another way is to rely on Go compiler and infer the datatype from the assigned value
	// card := "Ace of Spades"

	// Slice
	cards := deck{newCard(), newCard(), card, "Queen of Hearts"}
	cards = append(cards, "Six of Spades")

	cards.print()

	fmt.Println(card)
	fmt.Println(anotherCard)
}

// string here represents function's return type
func newCard() string {
	return "Five of Diamonds"
}
