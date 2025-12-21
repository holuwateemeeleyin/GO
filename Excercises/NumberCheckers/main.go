package main

import "fmt"

func main() {
	// Check if a number is postive, negative or zero
	var num int
	fmt.Println("Please enter a number: ")
	fmt.Scanf("%d", &num)
	if num > 0 {
		println("The number is positive")
	} else if num < 0 {
		println("The number is negative")
	} else {
		println("This is zero")
	}
}
