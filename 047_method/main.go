package main

import (
	"fmt"
)

type person struct {
	name       string
	profession string
}

func (x person) speak() {
	fmt.Printf("Hello! My name is %s, and I am a %s.\n", x.name, x.profession)
}

func main() {

	a := person{name: "Abebe", profession: "carpenter"}
	a.speak()

}
