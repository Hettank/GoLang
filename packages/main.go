package main

import (
	"fmt"
	"myproject/packages/math"
)

func main() {
	fmt.Println("Hello, World!")
	sum := math.Add(5, 3)
	fmt.Println("Sum:", sum)
	fmt.Println("Value of x:", math.X)
}
