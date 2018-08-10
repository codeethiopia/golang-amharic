package main

import (
	"fmt"
)

func main() {
	a := 8
	fmt.Printf("%b   %v\n", a, a)
	b := a << 2
	fmt.Printf("%b  %v\n", b, b)
	c := b >> 1
	fmt.Printf("%b  %v\n", c, c)
}
