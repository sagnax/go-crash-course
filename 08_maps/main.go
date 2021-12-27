package main

import "fmt"

func main() {
	// Define a map
	emails := make(map[string]string)

	// Assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	// declare map and values
	newEmails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	// add new kv
	newEmails["Mike"] = "mike@gmail.com"

	fmt.Println(newEmails)
	fmt.Println(newEmails["Sharon"])
}
