package main

// func add(a, b int) int {
// 	return a + b
// }

// func getlanguages() (string, string, bool) {
// 	return "javascript", "C++", true
//}

// FUNCTION INSIDE FUNCTION
func processit() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {

	// result := add(5, 5)
	// fmt.Println(result)

	// lang1, lang2, _ := getlanguages()
	// fmt.Println(lang1, lang2)

	fn := processit()
	fn(6)

}
