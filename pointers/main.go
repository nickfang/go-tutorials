package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// --or--
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// --auto assign-- string: "" int: 0 float: 0.0 bool: false
	// --or--
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "a@a.com",
			zipCode: 12345,
		},
	}
	// go will have updateFirstName coerce alex to pointer
	alex.updateFirstName("Alexis")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pPerson *person) updateFirstName(newFirstName string) {
	(*pPerson).firstName = newFirstName
}
