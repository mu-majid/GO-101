package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(byteSlice), ","))
}

func (d deck) shuffle() {
	// Random Seed everytime we run the app === random order of the cards
	source := rand.NewSource(time.Now().UnixNano()) // UnixNano is int64 - and this is our seed
	randomNumGenerator := rand.New(source)
	for i := range d {
		newPosition := randomNumGenerator.Intn(len(d) - 1)

		// swap in a fancy way
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
