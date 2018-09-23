package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i > 6 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("%d is even\n", i)
	}
}
