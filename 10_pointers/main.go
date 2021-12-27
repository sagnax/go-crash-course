package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// use * to read value from pointer
	fmt.Println(*b)

	// change value with pointer
	*b = 10
	fmt.Println(a)
}
