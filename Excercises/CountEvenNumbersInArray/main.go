package main

import "fmt"

func main() {
	//Look for even numbers in an array and add up
	arr := []int{5, 6, 8, 90, 64, 31, 65, 70, 0}
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			count++
		}
	}
	fmt.Println("The total count for Even numbers in the array is ", count)
}
