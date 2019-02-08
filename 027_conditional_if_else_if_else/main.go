package main

import "fmt"

func main() {
// if else
	x := 72

	if x%2 == 0 {
		fmt.Printf("%d is even\n", x)
	} else {
		fmt.Printf("%d is odd\n", x)
	}

// if else if else if ... else

	// Write a program that prints the numbers from 1 to 20. But for multiples of three print “Fizz” instead
	// of the number and for the multiples of five print “Buzz”. For numbers which are multiples of both three
	// and five print “FizzBuzz”.

	for i := 1; i < 21; i++ {
		if i%15 == 0 {
			fmt.Println(i,"FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println(i,"Buzz")
		} else if i%3 == 0 {
			fmt.Println(i,"Fizz")
		} else {
			continue
		}
	}


}
