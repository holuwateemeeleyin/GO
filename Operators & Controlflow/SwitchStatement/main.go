package main

import "fmt"

func main() {
	var score int = 10
	switch {
	case score >= 18:
		fmt.Println("The student passed quite well")
	case score >= 10:
		fmt.Println("The student averagely passed")
	case score <= 9:
		fmt.Println("The student failed ")
	default:
		fmt.Println("Invalid score ")
	}
}
