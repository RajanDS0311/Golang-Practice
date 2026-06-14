package main

import (
	"fmt"
)

// Slices are dynamic, flexible, and more powerful than arrays. They are built on top of arrays
// most useful construct in Go

func main() {

	// uninitialized slice are nil
	// 	var num []int
	// 	fmt.Println(num)        // Output: [] Default value for slice is nil
	// 	fmt.Println(num == nil) // Output: true
	// 	fmt.Println(len(num))   // Output: 0

	var num = make([]int, 1, 5) // make is used to create slice with specified length and capacity
	fmt.Println(num)            // Output: [0 0] Default values for int is 0
	fmt.Println(num == nil)     // Output: false
	fmt.Println(cap(num))

	// fmt.Println(cap(num)) // Output: 5  Capacity is the total number of elements that slice can
	//                                  hold without needing to allocate more memory.

	// num = append(num, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110) // append is used to add elements to the slice
	// fmt.Println(num)                                                // Output: [0 0 10 20 30 40]
	// fmt.Println(cap(num))                                           // Output: 10 Capacity is doubled when the slice exceeds its current capacity.

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// // copy function
	// copy(nums2, nums)

	// fmt.Println(nums, nums2)

	// slice operator
	// arr := []int{1, 2, 3, 4, 5}
	// fmt.Println(arr[0:2])
	// fmt.Print(arr[0:])

	// slice package
	// var num = []int{1, 2, 3, 4, 5}
	// var num1 = []int{1, 2, 3, 4, 5}
	// fmt.Println(slices.Equal(num, num1)) //equal compares elements

	// 2D slice
	// var don = [][]int{{7, 8, 9}, {4, 3, 2}}
	// fmt.Println(don)

}
