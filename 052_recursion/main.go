package main

import (
	"fmt"
)

func main() {
	fmt.Println("0! =", factorial(0))
	fmt.Println("1! =", factorial(1))
	fmt.Println("3! =", factorial(3))
	fmt.Println("10! =", factorial(10))
}

func factorial(x int) int {
	if x < 1 {
		return 1
	}
	return x * factorial(x-1)
}
