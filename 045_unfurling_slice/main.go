package main

import (
	"fmt"
)

func main() {

	l := []int{1, 3, 4, 5, 6, 7, 8, 9}
	m := sum(l...)
	fmt.Println(m)

	score("Abebe", l...)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func score(name string, s ...int) {
	totalScore := 0
	for _, v := range s {
		totalScore += v
	}
	fmt.Printf("%s scored : %d\n", name, totalScore)
}
