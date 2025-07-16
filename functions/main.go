package main

import "fmt"

func simpleFunction() {
	fmt.Println("this is simple function")
}

// function addtion and subtraction both do the same thing, just another way to write a function

// func subtraction(a, b int) int {
// 	return a - b
// }

func addition(a , b int) (result int) {
	result = a + b
	return
}

func main() {
	simpleFunction()
	additionVal := addition(5, 6)

	fmt.Println(additionVal)
}