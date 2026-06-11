package main

import "fmt"

// Arrays are fixed-size collections of elements of the same type. They are useful when you know the number of elements in advance and want to store them in a contiguous block of memory. In Go, arrays are defined with a specific size and type.
// numberes sequence of specified length and type
// In array is empty then, Int -> 0, String -> "" (empty string), Bool -> false

func main() {

	// var num [5]int
	// num[0] = 10
	// fmt.Println(num[0]) // Output: 10
	// fmt.Println(num)    // Output: [10 0 0 0 0] Default values for int is 0
	// array length
	// fmt.Println(len(num)) // Output: 5

	// var vals [4]bool
	// fmt.Println(vals) // Output: [false false false false] Default values for bool is false

	// var names [3]string
	// names[0] = "Alice"
	// fmt.Println(names) // Output: [Alice Bob Charlie]

	// To declare and initialize an array in one line.
	// names := [3]string{"Rajan", "ketan", "Pranay"}
	// fmt.Println(names)

	// 2D array
	matrix := [2][3]int{{1, 2, 3},
		{4, 5, 6}}
	fmt.Println(matrix) // Output: [[1 2 3] [4 5 6]]

}
