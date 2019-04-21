package main

import (
	"fmt"
)

func main() {
	x := []string{"banana", "orange", "grape", "cherry"}

	for index, value := range x {
		fmt.Println(index, value)
	}

	fmt.Println(x[2])
}
