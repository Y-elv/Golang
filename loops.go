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
	fmt.Println("\n")
	names := []string{"kizox", "rugira", "eddy", "rufido"}
	days := []string{"st", "nd", "rd", "th"}

	for i := 0; i < len(names); i++ {
		formatted := fmt.Sprintf("the %v %v name is %v\n", i+1, days[i], names[i])
		fmt.Print(formatted)
	}

	fmt.Println("\n")
	for index, value := range names {
		fmt.Printf("the value at index %v is %v\n", index, value)
	}
	//for ommiting one of this two (index,value)
	fmt.Println("\n")
	for _, value := range names {
		fmt.Printf("the value is %v\n", value)
	}
}
