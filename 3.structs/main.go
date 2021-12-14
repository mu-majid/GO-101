package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact:   contactInfo{email: "email", zipCode: 11122}}

	fmt.Println(alex)

	marc := person{"Marc", "Bill", contactInfo{email: "email", zipCode: 11122}}

	fmt.Println(marc)

	var julie person

	fmt.Println(julie)

	julie.firstName = "Julie"
	julie.lastName = "Doe"

	fmt.Println(julie)

}
