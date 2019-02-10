package main

import "fmt"

func main() {
	weather := "Cloudy"

	switch weather {
	case "Rain":
		fmt.Println("Take Umbrella")
	case "Sunny":
		fmt.Println("Wear shorts")
	case "Cold":
		fmt.Println("Wear Jacket")
	default:
		fmt.Println("do nothing")
	}
}
