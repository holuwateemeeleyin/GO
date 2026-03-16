package main

import "fmt"

// Display todo tasks
func displayTasks(tasks []string) {
	for index, task := range tasks {
		index++
		fmt.Printf("%d. %s \n", index, task)
	}
}

func main() {
	//Building a Todo app
	fmt.Println("These are the tasks in your Todo:")
	tasks := []string{"Apple", "Banana", "Orange"}
	displayTasks(tasks)
	fmt.Printf("You have %d items in the list \n", len(tasks))

}
