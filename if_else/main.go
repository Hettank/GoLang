package main

import (
	"fmt"
)

func main() {
	// if else statement
	if true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// if else if else statement
	if false {
		fmt.Println("false")
	} else if true {
		fmt.Println("true")
	} else {
		fmt.Println("else")
	}

	// switch case statement
	switch 1 {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	default:
		fmt.Println("default case")
	}


	var age int
	fmt.Println("Enter Age: ")
	fmt.Scan(&age)

	if age > 18 && age < 60{
		fmt.Println("You can drive")
	} else {
		fmt.Println("You cannot drive")
	} 
}