package main

import (
	"fmt"
)

func main() {
	s := "Selam, Tena Yistelegn"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	byteSlice := []byte(s)

	for i := 0; i < len(byteSlice); i++ {
		fmt.Printf("%#U ", byteSlice[i])
	}
	fmt.Println()
	for i := 0; i < len(byteSlice); i++ {
		fmt.Printf("%#x ", byteSlice[i])
	}	
	fmt.Println()
}