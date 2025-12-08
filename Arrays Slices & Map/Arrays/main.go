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

	//Multi dimension arrays:
	/* A multi-dimensional array is an array that has more than one direction
	Basically, it's an array of arrays
	Instead of storing data in a single line (1D array), it stores it in a grid-like structure.
	The most common types: 2D arrays (matrix/table), 3D arrays (cube), and higher dimensions.
	These are some examples of 2-dimensional array:
	arr := [3][2]int{{30, 40}, {50, 60}, {70, 80}}
	arr := [2][3]int{{1, 2, 3},{4, 5, 6},}

	These are some examples of 3-dimensional arrays:

	arr := [2][2][2]int{{{1,2},{3,4}},{{5,6},{7,8}}, is the same with
	arr := [2][2][2]int{
	{{1,2},{3,4}},
	{{5,6},{7,8}},
	}
	*/
}
