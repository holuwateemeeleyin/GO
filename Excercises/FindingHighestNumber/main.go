package main

import "fmt"

func main() {
	numbers := []int{3, 5, 9, 8, 0}
	highestNo := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > highestNo {
			highestNo = numbers[i]
		}
	}
	fmt.Println("The highest number is", highestNo)
}
