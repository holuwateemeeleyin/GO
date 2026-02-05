package main

import "fmt"

func declarePointer() {
	//to declare a pointer
	// var <pointer_name> *<data_type>
	var i *int
	var s *string
	fmt.Println(i)
	fmt.Println(s)
}

// Ways we can initialise a pointer
// 1. By including the data type
func initialisePointer1() {
	x := 10
	// To initialise a pointer
	// var <pointer_name> *<data_type> = &<variable_name>
	var ptr_x *int = &x
	// We are storing the address of i using the & operator sign

	fmt.Println("The address of x is:", ptr_x)
}

// 2. Not including the data type: In this case the data type will be determined by the compiler
// var <pointer_name> = &<variable_name>
func initialisePointer2() {
	s := "hello"
	var ptr_s = &s
	fmt.Println("The address of s is:", ptr_s)
	//If you are curious to know the datatype assigned to ptr_s
	fmt.Printf("The datatype assigned to ptr_s is: %T \n", ptr_s)

}

// 3. Using the shorthand operator
// we can initialise a pointer this way; <pointer_name> := &<variable_name>
func initialisePointer3() {
	s := "Hello world"
	ptr_s := &s
	fmt.Println("The address of ptr_s is:", ptr_s)
}
