package main

import (
	"fmt"
)

func printDetails(student string, subjects ...string) {
	fmt.Println("Hey, ", student, "here are your subjects - ")
	for _, sub := range subjects {
		fmt.Printf("%s, ", sub)
	}
}

// Anonymous Function
/*
 Anonymous function is a function that is declared without any named identifier to refer to it.
 They can accept inputs and return outputs just as standard functions.
 They can be used for containing functionality that need not be named and possibly for short term use
*/

func main() {
	printDetails("Sam", "Maths", "English")
	fmt.Println()
	num := 6
	//Recursive Function
	fmt.Println("The Factorial of ", num, "is ", factorial(num))

	// Using an Anonymous function
	x := func(a int, b int) int {
		return a * b
	}
	fmt.Println("First use of Anonymous function: ", x(20, 30))

	// Comment: Though the function is assigned to a variable x, the functionn does not have a name.
	// it can also be used in this format
	y := func(l int, b int) int {
		return l * b
	}(5, 10)
	fmt.Println("Second use: ", y)
	//Get User
	fmt.Println("This is Get User Information in getUser file")
	user, err := getUser(123)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("The user details: %+v", user)
}
