package main

import "fmt"

func main() {
	// A pointer is a variable that holds memory address of another variable.
	// They also point where the memory is allocated and provide ways to find or even change the value located at the memory location
	i := 10
	fmt.Printf("Pointer datatype and pointer address: %T, %v \n", &i, &i)
	// Derefrencing to get the dereferenced datatype and actual value
	fmt.Printf("The dereferenced data type and actual value: %T, %v \n", *(&i), *(&i))
	// or
	p := &i
	fmt.Printf("%T, %v \n", *p, *p)
}
