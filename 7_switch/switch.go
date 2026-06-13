package main

import (
	"fmt"
)

//in golang their is no need to use BREAK keyword

func main() {
	// simple switch statement
	// i := 10
	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// case 10:
	// 	fmt.Println("Ten")
	// default:
	// 	fmt.Println("Not One, Two or Three")
	// }

	// multiple conditions switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's the weekend")
	// default:
	// 	fmt.Println("It's a weekday")
	// }

	// TYPE SWITCH
	whoIam := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Its an integer")

		case string:
			fmt.Println("Its a string")

		case bool:
			fmt.Println("Its a boolean")

		default:
			fmt.Println("Others")
		}
	}

	whoIam("Hello, World!")
	whoIam(42)
	whoIam(true)
	whoIam(3.14)

}
