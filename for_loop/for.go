package main

import "fmt"

// for -> only construct in go for looping
func main() {
	// while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("================")

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
	
	fmt.Println("================")
	
	// infinite loop
	for {
		fmt.Println("infinite loop")
		break // break statement to exit the loop
	}
	
	fmt.Println("================")

	// for range loop
	for i := range 3 {
		fmt.Println(i) // prints 0, 1, 2, excluding 3
	}
}