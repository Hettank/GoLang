package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("What's your name?: ")
	// var name string
	// fmt.Scan(&name)

	// fmt.Println("\nHello", name)

	reader := bufio.NewReader(os.Stdin) // take user input
	name, _ := reader.ReadString('\n') // read the whole line until it hits the '\n' (new line)

	// fmt.Println("\nreader:", reader) // just a buffer string
	fmt.Println("\nname is :", name)
}

// A buffered reader is a type of reader that reads data from an underlying source. such as file or standard input (keyboard). and stores that data in a buffer. the buffer is a temporary storage area in a memory. Buffered readers are commonly used to improve the efficiency of input operations.