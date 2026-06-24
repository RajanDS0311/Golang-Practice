package main

import "fmt"

//Struct Embedding is a feature in Go where you put one struct inside another struct without giving it a field name.
//Struct embedding allows a struct to reuse the fields and methods of another struct.

//Struct embedding is similar to inheritance because it allows access to another struct's fields and methods.
// However, it is not true inheritance. Go uses composition (HAS-A relationship) instead of
// inheritance (IS-A relationship).

type parent struct {
	name string
	age  int
}

type child struct {
	hieght float32
	weight float32
	parent
}

func main() {
	// parent1 := parent{
	// 	name: "Rajan",
	// 	age:  21,
	// }

	child1 := child{
		hieght: 162,
		weight: 65,
		// parent: parent1,   OR
		parent: parent{ //inline add krna
			name: "Rajan",
			age:  21},
	}

	fmt.Println(child1)
	child1.parent.name = "Jay"
	fmt.Println(child1)
}
