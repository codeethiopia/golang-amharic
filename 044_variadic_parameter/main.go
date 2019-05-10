package main

import (
	"fmt"
)

func main() {

	x := sum(1)
	fmt.Println(x)

	y := sum()
	fmt.Println(y)

	z := sum(23, 454, 665, 776, 343, 444, 5657, 887)
	fmt.Println(z)

	l := []int{1, 3, 4, 5, 6, 7, 8, 9}
	m := sum(l...)

	fmt.Println(m)

}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
