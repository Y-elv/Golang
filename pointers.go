package main

import "fmt"

func updateName1(x *string) {
	*x = "angel"
}
func main() {
	name := "didier"
	n := &name
	// fmt.Println("the address of that variable name is :", n)
	// fmt.Println("the actual value is :", *n)
	fmt.Println(name)
	updateName1(n)
	fmt.Println(name)
}
