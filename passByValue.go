package main

import (
	"fmt"
)

// string  same to ints,bools,floats,arrays ,structs
func updateName(n string) string {
	x := n
	return x
}
func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}
func main() {
	name := "elvis"
	name = updateName("name")
	fmt.Println(name)
	menu := map[string]float64{
		"pie":       5.59,
		"ice cream": 3.99,
	}
	updateMenu(menu)
	fmt.Println(menu)
}
