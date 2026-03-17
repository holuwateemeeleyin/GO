package main

import "fmt"

// Display todo tasks
func displayTasks(tasks []string) {
	for index, task := range tasks {
		index++
		fmt.Printf("%d. %s \n", index, task)
	}
	fmt.Printf("You have %d items in the list \n", len(tasks))
}

// Add Task
func addTask(taskItems []string, newTask string) {
	updatedTasks := append(taskItems, newTask)
	fmt.Println("These are the updated tasks in your Todo:")
	displayTasks(updatedTasks)
}

// Delete Task
func removeTask(taskItems []string, delTask string) []string {
	var updatedTask []string
	for _, task := range taskItems {
		if task != delTask {
			updatedTask = append(updatedTask, task)
		}
	}
	fmt.Println("These are the updated tasks after removing", delTask)
	displayTasks(updatedTask)
	return updatedTask
}

func main() {
	//Building a Todo app
	fmt.Println("These are the tasks in your Todo:")
	tasks := []string{"Apple", "Banana", "Orange"}
	displayTasks(tasks)
	addTask(tasks, "Yam")
	removeTask(tasks, "Apple")
}
