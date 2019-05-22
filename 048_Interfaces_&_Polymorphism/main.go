package main

import (
	"fmt"
)

type dog struct {
	name  string
	color string
}

type cat struct {
	name  string
	color string
}

func (d dog) speak() string {
	return fmt.Sprintf("%s says, \"woof-woof\"\n", d.name)
}

func (c cat) speak() string {
	return fmt.Sprintf("%s says, \"meow\"\n", c.name)
}

type animal interface {
	speak() string
}

func myPet(a animal) {
	fmt.Println(a.speak())
}

func main() {
	d := dog{name: "Buchi", color: "Brown"}
	c := cat{name: "Wuro", color: "White"}

	myPet(d)
	myPet(c)
}
