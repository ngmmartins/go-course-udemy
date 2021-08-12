package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	ContactInfo
}

func main() {
	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		ContactInfo: ContactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.UpdateName("jimmy")
	jim.Print()
}

func (p Person) Print() {
	fmt.Printf("%+v", p)
}

func (p *Person) UpdateName(newFirstName string) {
	(*p).firstName = newFirstName
}
