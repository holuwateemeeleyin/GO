package main

import "fmt"

func main() {
	//Slices
	/* A slice is defined as the contnous segment of an underlying array
	A slice has 3 major components Pointer, length and Capacity
	Pointers are variables that hold memory addresses
	The length of a slice is the number of element it contains
	*/
	// Declaring a slice
	slice := []int{0, 48, 58, 47, 84, 4, 8}
	slices := slice[1:4]
	fmt.Println(slices)
	// It will print [48 58 47]
	fmt.Println("slice capacity:", cap(slice))
	//This will give 7
	fmt.Println("slices capacity:", cap(slices))
	// Slice capacity = number of elements from the slice’s start index to the end of the underlying array
	// That is, [1:4] will do length of slice(7) minus starting index of slices(1), IT WILL GIVE US 6
	//declaring and initialising slice with make
	fmt.Println("Declaring and initialising slice with make")
	slice = make([]int, 5, 8)
	fmt.Println("slice:", slice)
	fmt.Println("slice length:", len(slice))
	fmt.Println("slice capacity:", cap(slice))

}
