package main

import "fmt"

// Iterating over Data Structures
func main() {
	// RANGE ON SLICE
	// nums := []int{10, 20, 30, 40}
	// // sum := 0

	// for i, num := range nums {
	// 	// sum = sum + num
	// 	fmt.Println(i, num)

	// }

	//RANGE ON MAP
	// m := map[string]string{"Rajan": "Napur", "Pranav": "Pathri", "Abhay": "Gadchiroli"}
	// for k, v := range m { // 1st variable key ke liye and 2nd variable hai vo value ke liye
	// 	fmt.Println(k, v) // Agar sirf 1 hi variable lenge to sirf key hi print hoga
	// }

	//RANGE ON STRING
	//i is starting byte of rune
	// 300 tak -> 1 byte, uske baad 2 bytes
	for i, y := range "Golang" { //yaha par y particular alphabet ke liye
		// fmt.Println(i, y) // y will print unicode for each alphabet
		fmt.Println(i, string(y))
	}

}
