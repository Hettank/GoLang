// this is library package code
package math

// var x = 2 // this variable is not accessible to other packages
var X = 2 // this variable is accssible to other packages as well

func Add(a, b int) int {
	return a + b
}

// now i can use this package in my main.go file