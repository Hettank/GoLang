package main

import "fmt"

func acceptPointer(p *int) {
	fmt.Println("Value:", *p)
	fmt.Println("Address:", p)

	*p = 20
}

func main() {
	// val := 10
	// fmt.Println("Initial Value:", val)
	// acceptPointer(&val)
	// fmt.Println("Updated Value:", val)

	integer := 35
	ptr1 := &integer
	ptr2 := &ptr1

	fmt.Println("Value of integer:", **ptr2)
}