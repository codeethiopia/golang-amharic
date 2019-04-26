package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
	student   bool
}

func main() {
	p1 := person{
		firstName: "Abebe",
		lastName:  "Kebede",
		age:       40,
		student:   false,
	}

	fmt.Println(p1)

	p2 := person{
		firstName: "Yordanos",
		lastName:  "Negash",
		age:       25,
		student:   true,
	}

	fmt.Println(p2)

	fmt.Println(p1.firstName, p1.age)

	if p2.student {
		fmt.Println(p2.firstName, "is a student!")
	}

}
