package main

import "fmt"

// we only have for loop in golang, there is no while loop and do while loop.

func main() {
	// while loop
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	//infinite loop
	// for {
	// fmt.Println("Infinite Loop")
	// }

	// Classic for loop
	// for i := 1; i <= 5; i++ {

	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// using RANGE in for loop
	for i := range 11 {
		fmt.Println(i)
	}

}
