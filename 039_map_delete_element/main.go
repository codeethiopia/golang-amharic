package main

import (
	"fmt"
)

func main() {
	age := map[string]int{
		"Abebe": 20, "Zeritu": 21, "Ali": 25, "Kebede": 18}
	fmt.Println(age)

	delete(age, "Kebede")
	fmt.Println(age)
	// ማፕ ውስጥ ያልነበረ አባልን ቁልፍ በመጠቀም ለማስወገድ ብንሞክር ስህተት (error) አይመጣም ።
	delete(age, "Kassa")
	fmt.Println(age)

	// የምናስወግደው አባል ቁልፍ/ዋጋ በማፕ ውስጥ እንዳለና እንደሌለ  የ “comma ok idiom” በመጠቀም ማረጋገጥ እንችላለን
	if v, ok := age["Ali"]; ok {
		fmt.Println("value", v)
		delete(age, "Ali")
		fmt.Println(age)
	}
}
