package main

import (
	"fmt"
)

func main() {
	// equivalent to while loop
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println("\"While\" loop is Done!")

	for {
		fmt.Printf("Inside Continuous loop: %v\n",i)
		i++
	}
}
