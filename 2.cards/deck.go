package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings

// this line means deck extends (inherits) all functions/methods from 'Slice' type
type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, s := range cardsSuits {
		for _, v := range cardValues {
			cards = append(cards, s+" Of "+v)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
