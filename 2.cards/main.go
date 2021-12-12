package main

import "fmt"

func main() {
	var card string = "Ace of Spades"
	anotherCard := newCard()
	// Another way is to rely on Go compiler and infer the datatype from the assigned value
	// card := "Ace of Spades"

	// Slice
	cards := []string{newCard(), newCard(), card, "Queen of Hearts"}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	fmt.Println(card)
	fmt.Println(anotherCard)
	fmt.Println(cards)
}

// string here represents function's return type
func newCard() string {
	return "Five of Diamonds"
}
