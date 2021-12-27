package main

import (
	"fmt"
	"math"

	"github.com/sagnax/go-crash-course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("Test"))
}
