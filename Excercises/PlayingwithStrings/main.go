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

}
