package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	alex := person{firstName: "Alex", lastName: "Anderson"}

	fmt.Println(alex)

	marc := person{"Marc", "Bill"}

	fmt.Println(marc)

	var julie person

	fmt.Println(julie)

	julie.firstName = "Julie"
	julie.lastName = "Doe"

	fmt.Println(julie)

}
