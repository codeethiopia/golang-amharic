package main

import (
	"fmt"
)

func main() {
	x := []int{2, 4, 6, 8, 10, 12, 14, 16}

	fmt.Println(x)
	// የጠቋሚ ቦታን (index position) በመጠቀም ዋጋዎችን ማግኘት
	fmt.Println(x[2])
	// ስላይስን ለመቆራረጥ የምንጠቀመው የሁለት ነጥብ (colon) ኦፐሬተርን ነው
	fmt.Println(x[:])
	// ከሁለት ነጥቡ በስተግራ ያለው ፣ መነሻ ጠቋሚ ዋጋ ተቆርጦ የሚወጣው ስላይስ ውስጥ ይገባል
	fmt.Println(x[3:])
	// ከሁለት ነጥቡ በስተቀኝ ያለው ፣ መድረሻ ጠቋሚ ፣ ተቆርጦ የሚወጣው ስላይስ ውስጥ አይገባም
	fmt.Println(x[:3])
	// መነሻ እና መድረሻ ጠቋሚ
	fmt.Println(x[3:6])
}
