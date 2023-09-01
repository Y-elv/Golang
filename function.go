package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("good morning %v \n", n)

}
func sayBye(n string) {
	fmt.Printf("GoodBye %v \n", n)
}

func cycleName(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func Areaacircle(r float64) float64 {
	return math.Pi * r * r
}

func initials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var letter []string
	for _, v := range names {

		letter = append(letter, v[:1])
	}

	if len(letter) > 1 {
		return letter[0], letter[1]
	}
	return letter[0], "_"

}

func main2() {
	sayGreeting("elvis")
	sayBye("eddy")
	fmt.Println("\n")

	cycleName([]string{"dad", "mum", "brother"}, sayGreeting)
	fmt.Println("\n")
	cycleName([]string{"dad", "mum", "brother"}, sayBye)

	a1 := Areaacircle(10.5)
	a2 := Areaacircle(13.5)
	fmt.Println(a1, a2)
	fmt.Printf("%0.2f 0.1%f", a1, a2)

	//multiple returning values
	fmt.Println("\n")
	fmt.Println("\n")

	fn1, sn1 := initials("Mugisha Elvis")
	fmt.Println(fn1, sn1)

	fn2, sn2 := initials("Mugisha Elvis")
	fmt.Println(fn2, sn2)

	fn3, sn3 := initials("Mugisha")
	fmt.Println(fn3, sn3)

}
