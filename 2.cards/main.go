package main

import "fmt"

func main() {
	var card string = "Ace of Spades"
	// Another way is to rely on Go compiler and infer the datatype from the assigned value
	// card := "Ace of Spades"
	fmt.Println(card)
}
