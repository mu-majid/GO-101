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

* GO has two types of lists, namely, `Array` (Fixed length list of things), and `Slice` (An array that can grow or shrink)
* Elements' datatype in the array or slice should be the same.
* Note `sliceVar = append(sliceVar, "newElement")` will create a new slice and assign it to `sliceVar`, instead of modifying the old `sliceVar` slice.
* Note: Slices are passed by reference in Go.
* A slice is actually two data structure (under the hood in GO):
  1. An Array
  2. A Slice that keeps track of the array (holds -> ptr to head, capacity of array, length of the arr)

* Receiver functions are used to set a method on avriable of a custom type. Like type `deck`, we defined this receiver func on it.
```go
func (d deck) print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}
```

* to convert from a type to another -> typeWeWant(typeWeHave), example
```go
greeting := "Hi There!"

[]byte(greeting) // converting a string to Slice of byte
```