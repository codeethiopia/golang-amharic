package main

import (
	"fmt"
)

func main() {
	a := 123
	fmt.Printf("%T  %v\n", a, a)
	fmt.Printf("binary : %b  Decimal: %v\n", a, a)
	fmt.Printf("hex : %x  Decimal: %v\n", a, a)
	fmt.Printf("hex : %X  Decimal: %v\n", a, a)
	fmt.Printf("hex : %#x  Decimal: %v\n", a, a)
}
