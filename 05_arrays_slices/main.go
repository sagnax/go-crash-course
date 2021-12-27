package main

import "fmt"

func main() {
	// Arrays
	var fruits [2]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"

	var fruits2 = [2]string{"Apple", "Orange"}

	fruits3 := [2]string{"Apple", "Orange"}

	// Slices
	fruits4 := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruits, fruits2, fruits3, fruits4)
}
