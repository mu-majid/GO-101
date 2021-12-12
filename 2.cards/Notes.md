* Variables could be initialized outside of a function but cannot be assigned a value outside of the function.
``` go
package main

import "fmt"

var card string = "Ace of Spades" // WRONG

var card string   // RIGHT - Init
func main() {
	card = "Ace of Spades" // RIGHT - Assign
	fmt.Println(card)
}

```