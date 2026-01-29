package main

import "fmt"

func main() {
	//For loop
	var num int
	fmt.Printf("Enter the limit number: ")
	fmt.Scanf("%d", &num)
	for i := 1; i <= num; i++ {
		fmt.Printf("%d, ", i*i)
	}
}
