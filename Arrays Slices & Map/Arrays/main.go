package main

import "fmt"

func main() {
	/*An array is a collection of items stored in a single variable, where:
	1. All the items have the same data type
	2. The array has a fixed size (cannot grow or shrink)
	*/
	//In summary, Array has a fixed length, and elements should be of the same data type
	/* Why Arrays Are Useful:
	1. It stores multiple values using one name
	2. Easy to access values by index
	3. It makes looping easier
	4. Memory efficiency: It stores values for continous memory, making them efficient for fast reading, fast writing and low level operations
	5. It's also a foundation for other data structures like slices, queues, stacks, list and so on */

	fruits := [3]string{"apples", "orages", "banana"}
	fmt.Println(fruits)

	grades := [...]int{20, 30, 40, 50}
	fmt.Println(len(grades))

	//looping through array
	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}

	//Looping through an array with it's index
	fmt.Println("Looping through an array with it's index")
	for index, element := range grades {
		fmt.Println(index, "=>", element)
	}
}
