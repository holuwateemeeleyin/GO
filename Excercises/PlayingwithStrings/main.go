package main

import (
	"fmt"
	"strings"
)

func Swap(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := Swap("Hello", "World")
	fmt.Println("a, b= ", a, b)

	//Iterating Over Strings by Rune
	for index, runeValue := range "hello" {
		fmt.Printf("index: %d, value: %c \n", index, runeValue)
	}
	// String Operations
	fmt.Println("String Operations:")
	sentence := "Chelsea is the best football club in the world"
	fmt.Println("original Sentence: ", sentence)
	// 1. Splitting a string
	result := strings.Split(sentence, " ")
	fmt.Println("Splitted result:", result)
	// 2. Replacing a string
	replacedString := strings.Replace(sentence, "Chelsea", "Arsenal", 1)
	fmt.Println("Replaced String:", replacedString)
	//3. Upper and Lower case
	fmt.Println("UpperCase:", strings.ToUpper(sentence))
	fmt.Println("LowerCase:", strings.ToLower(sentence))
	//4. Trim
	fmt.Println("Trim:", strings.Trim("  Trim     ", " "))
	//Subset
	fmt.Println("Subset:", sentence[0:19])

}
