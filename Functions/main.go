package main

import "fmt"

func printDetails(student string, subjects ...string) {
	fmt.Println("Hey, ", student, "here are your subjects - ")
	for _, sub := range subjects {
		fmt.Printf("%s, ", sub)
	}
}

// Recursive function
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	printDetails("Sam", "Maths", "English")
	fmt.Println()
	num := 6
	fmt.Println("The Factorial of ", num, "is ", factorial(num))
}
