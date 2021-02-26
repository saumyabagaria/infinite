package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	address   string
}

func main() {
	var emp1 person
	emp1 = person{firstName: "Anand", lastName: "Babu", age: 20, address: "Bangalore"}
	fmt.Printf(" Details %q", emp1)
}
