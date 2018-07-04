package main

import (
	"fmt"
)

var a bool

func main() {
	a = true
	b := 2
	c := 3
	fmt.Println(a)
	fmt.Println(b == c)
}