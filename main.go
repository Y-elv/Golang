package main

import "fmt"

func main() {

	var name string = "elvis"
	secondName := "angel"
	var age int = 0
	var age1 uint = 2

	fmt.Println("hello", name, "your age is:", age)
	fmt.Print(secondName, age1, "\n")
	fmt.Printf("my name is %v and my age is %v", name, age)
}
