package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// 4 ና 5 ን ለማስወገድ
	x = append(x[:3], x[5:]...)
	fmt.Println(x)
}
