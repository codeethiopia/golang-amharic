package main

import (
	"fmt"
)

var a int

type እንጀራ int // እንጀራ የሚባል አዲስ “type”

var b እንጀራ // አይነቱ “እንጀራ” የሆነ ተባራይ ታወጀ (a variable of type "እንጀራ" declared)

func main() {
	a = 42
	b = 43
	// a = b
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
