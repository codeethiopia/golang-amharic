package main

import (
	"fmt"
)

func main() {
	func() { fmt.Println("anonymous function") }()
}
