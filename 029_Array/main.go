package main

import (
	"fmt"
)

func main() {
	var x [5]int
	y := [4]int{}
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x[4] = 10
	fmt.Println(x)
	fmt.Printf("%T\n",y)
}
