package main

import (
	"fmt"
	"maps"
)

// Map -> map is collection of key-value pairs where each key is unique,unordered data structure ,(Dictionary).
func main() {

	// creating a map using make function
	// we use make function when we don't know the elements beforehand
	// m := make(map[string]string)

	// adding key-value pairs to the map
	// m["name"] /*key */ = "John" /* value */
	// m["sex"] /*key */ = "male"  /* value */

	// get the element from the map
	// fmt.Println(m["name"], m["sex"]) // Output: John male

	// if key is not present in the map, it returns the zero value of the value type
	// fmt.Println(m["age"]) // Output: "" zero value for string is empty string

	// m := make(map[string]int)
	// m["age"] = 21
	// m["height"] = 162
	// m["weight"] = 62
	// fmt.Println(m["age"])    // Output: 21
	// fmt.Println(m["height"]) // Output: 0 zero value for int is 0

	// fmt.Println(len(m))
	// fmt.Println(m)

	// delete a key-value pair from the map
	// delete(m, "weight")
	// fmt.Println(m) // Output: map[age:21 height:162]

	// // all clear the map
	// clear(m)
	// fmt.Println(m) // Output: map[]

	// creating a map without using make function
	// we use this when elemnts are available to us from starting
	// m := map[string]int{
	// 	"age":    21,
	// 	"height": 162,
	// 	"weight": 62,
	// }
	// fmt.Println(m) // Output: map[age:21 height:162 weight:62]

	// getting element from the map using key
	m := map[string]int{
		"size":   10,
		"mobile": 2098239839,
	}
	value, ok := m["size"]       // ok is a boolean variable that indicates whether the key is present in the map or not
	fmt.Println("Value:", value) // Output: Value: 10
	if ok {
		fmt.Println("all ok") // Output: Value: 10
	} else {
		fmt.Println("Key not found")
	}

	// check if two maps are equal
	m1 := map[string]int{
		"size":   10,
		"mobile": 2098239839}

	m2 := map[string]int{
		"size":   10,
		"mobile": 2098239839}

	fmt.Println(maps.Equal(m1, m2))
}
