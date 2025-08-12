package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Product struct {
	Name string
	Price float64
}

func (p Person) Greet() string {
	// Method that belongs to the Person type
	return "Hello, my name is " + p.Name
}

func MethodGreet(p Person) int {
	// Regular function that takes a Person as an argument
	return p.Age
}

func main() {
	// Methods in Go: Go provides us a feature methods which allows us to assign functions to a struct type.
	// Any types can have methods, including built-in types like int, string, etc.

	newPerson := Person{Name: "Het", Age: 21}
	anotherPerson := Person{Name: "Maya", Age: 22}

	fmt.Println(anotherPerson.Greet())
	fmt.Print(newPerson.Greet())

	ageExample := Person{Name: "Example", Age: 30}
	fmt.Println("\n",MethodGreet(ageExample))

	// A method is attached to a type (like Person), and only that type (and its aliases) can use it with the dot notation.
	// the below instance will not work as Product does not have a Greet method.
	// prod := Product{Name: "Laptop", Price: 999.99}
	// prod.Greet()
}