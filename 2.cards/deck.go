package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings

// this line means deck extends (inherits) all functions/methods from 'Slice' type
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
