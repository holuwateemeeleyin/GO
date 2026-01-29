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
	fmt.Println("2-dimensional array")
	arr := [3][2]int{{30, 40}, {50, 60}, {70, 80}}
	fmt.Println(arr[2][1])
	//Looping a 2-dimensional array
	fmt.Println("Looping 2-dimensional array")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	// Todo: declare a 3-dimensional array and loop through it
	fmt.Println("Looping 3-dimensional array")
	arr3 := [2][3][2]int{
		{{1, 2}, {3, 4}, {5, 6}},
		{{7, 8}, {9, 10}, {11, 12}},
	}
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Print(arr3[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println("Another 3-dimensional array")
	arr4 := [2][3][3]int{
		{{1, 2, 13}, {3, 4, 14}, {5, 6, 15}},
		{{7, 8, 16}, {9, 10, 17}, {11, 12, 18}},
	}
	for i := 0; i < len(arr4); i++ {
		for j := 0; j < len(arr4[i]); j++ {
			for k := 0; k < len(arr4[j]); k++ {
				fmt.Print(arr4[i][j][k], " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}

}
