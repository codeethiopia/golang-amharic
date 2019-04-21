package main

import (
	"fmt"
)

func main() {
	x := []string{"banana", "orange", "grape", "cherry"}

	for index, value := range x {
		fmt.Println(index, value)
	}
	// የጠቋሚ ቦታን (index position) በመጠቀም
	fmt.Println(x[2])
}
