package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num * 5
}

func main() {
	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num

	num := 25
	ptr := &num

	// default value of any pointer is nil

	// fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", ptr)
	fmt.Println("Value at address of num:", *ptr) // dereferencing the pointer

	value := 10
	modifyValueByReference(&value)
	fmt.Println("Value after modification:", value)
}