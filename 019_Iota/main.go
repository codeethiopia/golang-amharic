package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)

const (
	d = iota
	e
	f
	g
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
