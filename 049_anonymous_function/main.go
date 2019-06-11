package main

import (
	"fmt"
)

func main() {
	func() { fmt.Println("anonymous function executed") }()

	func(x int, y int) {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}(2, 3)

}
