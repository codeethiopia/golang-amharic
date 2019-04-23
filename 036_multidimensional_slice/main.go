package main

import (
	"fmt"
)

func main() {
	fruits := []string{"grape", "cherry", "orange", "lime"}
	animals := []string{"dog", "horse", "cat"}
	x := [][]string{fruits, animals}
	fmt.Println(x)
}
