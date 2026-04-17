package main

import (
	"fmt"
	"strings"
)

func countVowels(input string) int {
	count := 0
	input = strings.ToLower(input)

	for _, char := range input {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			count++
		}
	}
	return count
}

func main() {
	var text string

	fmt.Print("Enter a string: ")
	fmt.Scanln(&text)

	result := countVowels(text)
	fmt.Println("Number of vowels:", result)
}