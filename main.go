package main

import "fmt"

func main() {

	var name string = "elvis"
	secondName := "angel"
	var age int = 0
	var age1 uint = 2
	score := 30.2

	fmt.Println("hello", name, "your age is:", age)
	fmt.Print(secondName, age1, "\n")
	fmt.Printf("my name is %v and my age is %v \n", name, age)
	fmt.Printf("the type of this age is %T \n", age)
	fmt.Printf("you scored %f points!\n", score)
	var str = fmt.Sprintf("my name is %v and my age is %v \n", name, age)
	println(str)

}
