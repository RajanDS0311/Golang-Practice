package main

import "fmt"

func main() {
	// age := 5

	// if age > 18 {
	// 	fmt.Println("You are an adult.")
	// } else if age >= 13 && age <= 18 { // and -> &&, or -> ||, not -> ! OPERATORS
	// 	fmt.Println("You are a teenager.")
	// } else {
	// 	fmt.Println("You are a kid.")
	// }

	// we can also declare a variable in if statement like this
	if age := 5; age > 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 && age <= 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a kid.")
	}

}
