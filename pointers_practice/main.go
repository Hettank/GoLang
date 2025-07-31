package main

import "fmt"

func someFunction(number *int) {
	*number += 6
}

func main() {
	var number int
	number = 78

	var pointer *int
	pointer = &number

	fmt.Println("Address of number:", pointer)
	fmt.Println("Value at address of number:", *pointer)

	someFunction(&number)
	fmt.Println("Value after modification:", number)
}