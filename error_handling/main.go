package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator must not be 0")
	}

	return a / b, nil
}

func main() {
	dividedVal, _ := divide(5, 0)

	fmt.Println("\nval:", dividedVal)
}

// in go, the underscore is used as a blank identifier, it serves as a write-only variable that allows you to discard values returned by a function or to ignore specific values when you are not intrested in using them