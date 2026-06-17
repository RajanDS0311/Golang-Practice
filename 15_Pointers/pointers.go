package main

import "fmt"

//pass by Value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

//pass by reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)

}

func main() {
	num := 1

	// changeNum(num)
	changeNum(&num)

	// fmt.Println("Memory Adress of num = 1 is", &num) //& use krte hai address k liye
	fmt.Println("After chanheNum in main", num)

}
