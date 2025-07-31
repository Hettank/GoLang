package main

import "fmt"



func simpleFunction() {
	fmt.Println("this is simple function")
}

// function addtion and subtraction both do the same thing, just another way to write a function

// func subtraction(a, b int) int {
// 	return a - b
// }

func compute(fn func(int, int) int) int {
	return fn(10, 6)
}

func addition(a , b int) (result int) {
	result = a + b
	return
}

func swap(x, y string) (string, string) {
	return y, x // function can return multiple values
}



func main() {
	assigned := func() string {
		return "this is the function assigned to a variable"
	}

	fmt.Println("from compute: ", compute(addition))

	simpleFunction()
	additionVal := addition(5, 6)

	fmt.Println(additionVal)

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// v1 is of type int64 and v2 is of type int32, even though they have the same value and same type int, but still you cannot assign v1 to v2 directly because the internal data types are different.
	// var v1 int64
	// var v2 int32
	// v1 = 32
	// v2 = 32

	// v2 = v1



	fmt.Println(assigned())
}