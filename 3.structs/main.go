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

	alex.print()

	marc := person{"Marc", "Bill", contactInfo{email: "email", zipCode: 11122}}

	marc.print()

	marcPointer := &marc
	marcPointer.updateFname("Marco")
	marc.print()

	var julie person

	julie.print()

	julie.firstName = "Julie"
	julie.lastName = "Doe"

	julie.print()

}

func (pointerToPerson *person) updateFname(newFname string) {
	(*pointerToPerson).firstName = newFname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
