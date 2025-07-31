package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Hello,world!"
	parts := strings.Split(data, ",")

	for i, part := range parts {
		fmt.Println("Part", i, ":", part)
	}
}