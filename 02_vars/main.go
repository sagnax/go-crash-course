package main

import "fmt"

func main() {
	// MainTypes
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// types can be inferred

	// Using var or const or the shorthand :=
	var name string = "Lucas"
	var age = 27
	const isCool = true
	size := 1.3
	var1, var2 := "bata", "ta"

	fmt.Println(name, age, isCool, size, var1, var2)
	fmt.Printf("%T\n", name)
}
