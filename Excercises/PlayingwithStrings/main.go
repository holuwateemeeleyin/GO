package main

import "fmt"

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
}
