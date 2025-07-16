package main

import "fmt"

// Slices are dynamic data structures that provides a more powerful alternative to arrays, no need to explicitly define length

func main() {
	numbers := []int {1, 2, 3, 4, 5}
	fmt.Println("Before appending")
	fmt.Println("Numbers : ", numbers)
	fmt.Printf("Numbers has data of type : %T\n", numbers)

	numbers = append(numbers, 4, 2, 20, 19)

	fmt.Println("After appending")
	fmt.Println("Numbers : ", numbers)

	newSlice := make([]int, 4, 5) // => 4 is length and 5 is capacity 
	// capacity: capacity is like the maximun number of data you can store.
	// four 0's would be stored inside of this slice cuz we have given the length of 4 and by default it will initialize it with 0
	fmt.Println("\nSclice with make function: ", newSlice)
}