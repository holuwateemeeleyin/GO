package main

import "fmt"

func main() {
	var fruit string
	fmt.Printf("Enter the name of fruit:")
	fmt.Scanf("%s", &fruit)
	if fruit == "apple" {
		fmt.Println("Apple is a fruit")
	} else {
		fmt.Println("This is not an apple")
	}
}
