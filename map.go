package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"toffe pudding": 3.55,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//looping
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phoneNumber := map[int]string{
		37888: "elvis",
		56557: "eddy",
		44444: "anel",
	}
	fmt.Println(phoneNumber)
	fmt.Println(phoneNumber[44444])
	phoneNumber[44444] = "ian"
	fmt.Println(phoneNumber)
}
