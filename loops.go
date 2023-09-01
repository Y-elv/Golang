package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("the value of x is :", x)
		x++

	}
	fmt.Println("\n")

	for i := 0; i < 5; i++ {
		fmt.Println("the value of i is :", i)
	}

	names := []string{"kizox", "rugira", "eddy", "rufido"}
	days := []string{"st", "nd", "rd", "th"}

	for i := 0; i < len(names); i++ {
		formatted := fmt.Sprintf("the %v %v name is %v\n", i+1, days[i], names[i])
		fmt.Print(formatted)
	}
}
