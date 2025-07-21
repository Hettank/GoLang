// A map is a datastructure that stores key-value pairs. It is similar to a dictionary in Python or an object in JavaScript. Maps are used to store data in a way that allows for fast retrieval based on keys.
package main

import "fmt"

func main() {
	// name <-> age
	people := make(map[string]int)
	people["het"] = 21
	people["meet"] = 20
	people["jash"] = 20

	fmt.Println(people["het"])

	age, ok := people["meeet"]

	if ok {
		fmt.Println("Meet's age is", age)
	} else {
		fmt.Println("Meet not found")
	}

	// traversing a map
	for i := range people {
		fmt.Println("Name:", i, "Age:", people[i])		
	}

	// Initializing a map with values at the time of declaration
	newMap := map[string]int {
		"Ahmedabad": 25,
		"Katch": 30,
		"Baroda": 25,
	}

	fmt.Println("\n")
	for city, temprature := range newMap {
		fmt.Println("City:", city, "Temprature:", temprature)
	}
}