package main

import "fmt"

func main() {
	var name [5]string

	name[0] = "Het"
	name[1] = "Tank"

	fmt.Println("array:", name)
	
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("array:", numbers)
	fmt.Println("length of number are:", len(numbers))
}