package main

import "fmt"

func main() {
	// type infer
	var name = "hello"
	var age int = 21

	// shorthand syntax
	newName := "world"

	const (
		pi = 3.14
		age2 = 22
	)

	fmt.Println(pi)
	fmt.Println(age2)
	
	fmt.Print(name)
	fmt.Println(age)
	fmt.Println(name, newName)	

	var nums [5]int
	fmt.Print(len(nums))
}