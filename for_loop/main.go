package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	numbers := []int {10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index of Numbers is %d, and value is %d\n", index, value)
	}
}