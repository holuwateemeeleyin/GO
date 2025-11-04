package main

import "fmt"

// Knowing how to read data inputted in the terminal
func main() {
	/*
		//Using Scanf
		var name string
		fmt.Print("Please input your name: ")
		fmt.Scanf("%s", &name)
		fmt.Println("Welcome,", name)*/

	//The fmt.ScanF returns two values, which are:
	//1. Count(n): The number of successful scanned Items, and
	//2. err: An error that indicates what went wrong
	//For example:
	var a string
	var b int
	fmt.Print("Enter a string and an integer: ")
	count, err := fmt.Scanf("%s %d", &a, &b)

	//Demonstrating how scanf works
	fmt.Println("Count:", count) //This returns the count or number of the scanned Items that are successful
	fmt.Println("error:", err)   //If there is an error, this returns it
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	//if you input Timi 12, output will be:
	/* Output will be:
	Count: 2
	error: <nil>
	a: Tim
	b: 25 */
}
