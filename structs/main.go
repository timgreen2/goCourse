package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	// jim := person{"jim", "Anderson"} // go interprets the order of inputs as the order that they are declared in the person type
	jim := person{firstName: "jim", lastName: "Anderson", contactInfo: contactInfo{email: "jim@email.com", zip: 12345}}

	jim.updateName("jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
