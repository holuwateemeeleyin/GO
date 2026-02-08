package main

import "fmt"

// A struct is a collection of fields. It is used to group data together to form records. It is a user-defined data type that allows us to create complex data structures.
// We can define a struct using the type keyword followed by the struct name and the struct fields enclosed in curly braces.
type Person struct {
	Name   string
	Colour string
	Age    int
}

func main() {
	// To create an instance of a struct, we can use the struct literal syntax. We can also use the new keyword to create a pointer to a struct.
	p := Person{
		Name:   "Timi",
		Colour: "Green",
		Age:    29,
	}

	//user3 := new(User)
	// To access the fields of a struct, we can use the dot notation. We can also use the pointer to access the fields of a struct.
	fmt.Printf("User %+v\n", p)
}
