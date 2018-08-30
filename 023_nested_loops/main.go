package main

import (
	"fmt"
)

func main() {
	//ማስነሻ ;መወሰኛ ; ማስቀጠያ
	for i := 0; i < 11; i++ {
		fmt.Printf("Outer loop : %v\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\tInner loop : %v\n", j)
		}
	}
}
