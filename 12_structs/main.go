package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// greeting method (value receiver)
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and i am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer recevier)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	person1 := Person{firstName: "Sam", lastName: "Smith", city: "Florida", gender: "F", age: 20}
	// alternative
	person2 := Person{"Bob", "Nevar", "Florida", "M", 23}

	fmt.Println(person1)
	//fmt.Println(person2)
	fmt.Println(person1.firstName)
	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())
	person2.getMarried("Thompson")
	fmt.Println(person2.greet())
}
