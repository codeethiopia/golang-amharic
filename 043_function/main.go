package main

import (
	"fmt"
)

func main() {
morning()
afternoon("Abebe")
selamta := evening("Abebe")
fmt.Println(selamta)

x,y := holiday("Abebe","Zeritu")

fmt.Println(x,y)
}

func morning() {
	fmt.Println("Endemin Aderachuh")
}

func afternoon(name string) {
	fmt.Printf("Tena Yistelegn, %s !\n",name)
}

func evening(name string) string {
	return fmt.Sprintf("Dehna eder, %s!\n",name)
}

func holiday(name1 string,name2 string) (string,string) {
	a := fmt.Sprintf("Enkuan Adereseh, %s!\n",name1)
	b := fmt.Sprintf("Enkuan Aderesesh, %s!\n",name2)
	return a,b
}

