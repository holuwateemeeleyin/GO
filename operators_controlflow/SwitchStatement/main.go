package main

import "fmt"

// Using a switch statement, determine the score remark of students who attempted a CA test of 10marks
func main() {
	var score int
	fmt.Printf("Enter Student Score: ")
	fmt.Scanf("%d", &score)
	switch {
	case score >= 16 && score <= 20:
		fmt.Println("The student passed quite well")
	case score >= 10 && score < 16:
		fmt.Println("The student averagely passed")
	case score <= 9:
		fmt.Println("The student failed ")
	default:
		fmt.Println("Invalid score ")
	}
}
