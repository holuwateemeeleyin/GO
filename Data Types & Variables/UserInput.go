package main
import "fmt"

//Knowing how to read data inputted in the terminal
func main()  {
	//var name string
	//fmt.Print("Please input your name: ")
	//fmt.Scanf("%s", &name)
	//fmt.Println("Welcome,", name)

	//The fmt.ScanF returns two values, which are:
	//1. Count(n): The number of successful scanned Items
	//2. err: An error that indicates what went wrong
	//For example:
	var a string
	var b int
	fmt.Print("Enter a string and an integer: ")
	count, err := fmt.Scanf("%s", "%d", &a, &b)

	//Demonstrating how scanf works
	fmt.Println("Count:", count)
	fmt.Println("error:", err)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}