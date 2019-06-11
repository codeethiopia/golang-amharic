package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Println("func expression")
	}
	a()
}
