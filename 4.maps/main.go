package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colors)

	var otherColors map[string]string

	fmt.Println(otherColors)

	otherColors2 := make(map[string]string)

	otherColors2["white"] = "#ffffff"

	fmt.Println(otherColors2)

	delete(otherColors2, "white")

	fmt.Println(otherColors2)

	for color, hex := range colors {
		fmt.Println(color + " : " + hex)
	}
}
