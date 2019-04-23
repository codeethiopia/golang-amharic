package main

import (
	"fmt"
)

func main() {
	x := make([]int, 3, 5)
	fmt.Println(x)
	fmt.Println(len(x), cap(x))
	
	x[2] = 23
	fmt.Println(x)
	fmt.Println(len(x), cap(x))
	
	x = append(x, 34)
	fmt.Println(x)
	fmt.Println(len(x), cap(x))

	x = append(x, 45)
	fmt.Println(x)
	fmt.Println(len(x), cap(x))

	x = append(x, 56)
	fmt.Println(x)
	fmt.Println(len(x), cap(x))

}
