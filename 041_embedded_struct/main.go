package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
}
// አንድ ስትረክት በሌላ ስትረክት ውስጥ ሊከተት ይችላል
type student struct {
	person
	school string
	grade  float32
}

func main() {
	p1 := person{
		firstName: "Abebe",
		lastName:  "Kebede",
		age:       20,
	}

	fmt.Println(p1)

	p2 := person{
		firstName: "Yordanos",
		lastName:  "Negash",
		age:       25,
	}

	s1 := student{
		person: p1,
		school: "AAU",
		grade:  3.9,
	}

	s2 := student{
		person: p2,
		school: "AAU",
		grade: 4.0,
	}

	fmt.Println(s1)

	fmt.Println(s2)
	// በውስጠኛው ስትረክት ውስጥ ያሉትን አይነቶች በቀጥታ ማግኘት እንችላለን ( the inner type gets promoted to the outer type)
	fmt.Println(s1.firstName,s1.age,s1.school)
}
