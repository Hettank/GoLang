package main

import (
	"fmt"
	"time"
)

func main() {
	currntTime := time.Now()
	fmt.Println("Current Time:", currntTime)
	fmt.Println("Current Time:", currntTime.Format("2006-01-02 15:04:05"))
}