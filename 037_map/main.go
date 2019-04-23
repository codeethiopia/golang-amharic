package main

import (
	"fmt"
)

func main() {
	// የቁልፍ እና ዋጋ ማከማቻ
	age := map[string]int{
		"Abebe": 20, "Zeritu": 21, "Ali": 25, "Kebede": 18}
	fmt.Println(age)

	// ማፕ በፍጥነት መገኘት ያለበትን ውህብ ለማስቀመጥ ምቹ ነው ።
	fmt.Println(age["Kebede"])

	// ማፕ ውስጥ ያለና የሌለ ዋጋን በቀላሉ መለየት ይቻላል
	x, ok := age["Abebe"]
	fmt.Println(x, ok)

	y, ok := age["Chala"]
	fmt.Println(y, ok)

	// ማፕ ውስጥ ያሉ የቁልፍ-ዋጋ ጥንዶች ቅደም ተከተል የላቸውም ። ስለዚህ በማፕ ውስጥ ያሉ ቁልፎችንና ዋጋዎችን
	//  ብናትም (print) የሚወጡበት ቅደም ተከተል ያልተወሰነ (random) ነው ።
	for x := range age {
		fmt.Println(x)
	}
}
