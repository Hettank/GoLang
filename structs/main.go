package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Address struct {
	street string
	city   string
}

type Contact struct {
	phone   string
	email   string
}

type Employee struct {
	PersonDetails Person
	AddressDetails Address
	ContactDetails Contact
}

func main() {
	// het.name = "Het"
	// het.age = 30

	// fmt.Println(het)

	het := Person {
		name: "Het",
		age: 21,
	}

	aanchal := Person {
		name: "Aanchal",
		age: 24,
	}

	fmt.Println("Person 1:", het)
	fmt.Println("Person 2:", aanchal)

	employee := Employee {
		PersonDetails: Person{
			name: "Aanchal",
			age: 24,
		},
		AddressDetails: Address{
			street: "Vaishno Devi Circle",
			city:  "Ahmedabad",
		},
		ContactDetails: Contact{
			phone: "9428936537",
			email: "aanchal@gmail.com",
		},
	}

	fmt.Println("\nEmployee Details:", employee)
}