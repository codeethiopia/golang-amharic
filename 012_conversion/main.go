package main

import (
	"fmt"
)

var a int

type age int // new type age

type ርዝመት int // ርዝመት የሚባል አዲስ “type”

var b ርዝመት // አይነቱ “ርዝመት” የሆነ ተባራይ ታወጀ (a variable of type "ርዝመት" declared)

var c age // variable of type "age" declared

func main() {
	a = 4
	b = ርዝመት(a)
	c = age(b)
	// a = b
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Printf("%v\t%T\n", c, c)
}