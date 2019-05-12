package main

import (
	"fmt"
)

func main() {
	defer bye()
	greeting()
}

func greeting() {
	fmt.Println("Good Morning")
}

func bye() {
	fmt.Println("Bye")
}
