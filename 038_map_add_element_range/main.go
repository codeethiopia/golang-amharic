package main

import (
	"fmt"
)

func main() {
	age := map[string]int{
		"Abebe": 20, "Zeritu": 21, "Ali": 25, "Kebede": 18}
	fmt.Println(age)

	// አዲስ አባል መጨመር
	age["Kassa"] = 24
	fmt.Println(age)

	// rangeን በመጠቀም ማፕ ውስጥ ያሉትን ቁልፍና ዋጋዎች ማግኘት ይቻላል
	for k, v := range age {
		fmt.Println(k, v)
	}
}
