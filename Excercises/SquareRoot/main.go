package main

import (
	"errors"
	"fmt"
	"math"
)

// Finding square root of a number
func SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("There is no square root of a negative value")
	}
	return math.Sqrt(x), nil
}
func main() {
	result, err := SquareRoot(-25)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Square root:", result)
}
