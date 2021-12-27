package main

import "fmt"

func main() {
	ids := []int{12, 42, 43, 65, 14, 23}

	// loop through ids
	for i, ids := range ids {
		fmt.Printf("%d - ID: %d\n", i, ids)
	}
	// not using index
	for _, ids := range ids {
		fmt.Printf("ID: %d\n", ids)
	}
	// add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// range with maps
	newEmails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	for k, v := range newEmails {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k := range newEmails {
		fmt.Printf("Name: %s\n", k)
	}
	for _, v := range newEmails {
		fmt.Printf("Email: %s\n", v)
	}
}
