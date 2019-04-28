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

	// ስም-አልባ የምንለው ስትረክት ምንም ገላጭ ስያሜ የሌለው ነው ።
	p2 := struct {
		firstName string
		lastName  string
		age       int
		student   bool
	}{
		firstName: "Yordanos",
		lastName:  "Negash",
		age:       25,
		student:   true,
	}

	fmt.Println(p2)

	fmt.Println(p2.firstName, p2.age)

}
