package main

import (
	"fmt"
)

func main() {
	a := foo("Abebe beso bela")
	b := a()
	fmt.Println(b)
	fmt.Printf("The return type of func foo is : %T\n", a)
}

func foo(x string) func() string {
	return func() string {
		return x
	}
}
