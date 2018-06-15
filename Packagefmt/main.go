package main

import (
	"fmt"
)

var a int
var b string = "አበበ በቂላ"
var c bool
var d bool = true

func main() {
	e := 42
	f := "አበበ በሶ በላ"
	g := `ቀስ በቀስ እንቁላል በእግሩ ይሄዳል`
	h := `Paulos Gnogno was an extraordinary Ethiopian journalist, historian and author who was passionate
	about informing and educating Ethiopians about a wide range of issues.`
	i := 1387473.384859438
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%t\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%d\n", e)
	fmt.Printf("%s\n", f)
	fmt.Printf("%s\n", g)
	fmt.Printf("%v\n", h)
	fmt.Printf("%.3f\n", i)
}
