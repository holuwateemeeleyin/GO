package main
import "fmt"

//Knowing how to read data inputted in the terminal
func main()  {
	var name string
	fmt.Print("Please input your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Welcome, " name)
}