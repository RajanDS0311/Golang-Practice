package main

import "fmt"

const age = 21 //global constant variable and can be accessed anywhere in program

func main() {
	const name = "Golang"
	const age = 23 //local constant variable with same name as global var hides it and overlaps

	// also can write like this i.e,.....CONSTANT GROUPING
	const (
		Rajan = 21
		saal  = 2005
	)
	fmt.Println(age)
	fmt.Println(saal)
	fmt.Println(name)
	fmt.Println(Rajan)
}
