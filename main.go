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
	var marks = [3]int{10, 11, 23}
	fmt.Println(marks, len(marks))
	fmt.Println(len(marks))

	fmt.Println("after modification")
	marks[1] = 12
	fmt.Println(marks, len(marks))

	names := [2]string{"hey", "you"}
	fmt.Println(names)

	//slices

	var scores = []int{1, 2, 3}
	scores[2] = 4

	scores = append(scores, 35)
	fmt.Println(scores, len(scores))

	//slice range

	rangeOne := scores[1:3]
	fmt.Println(rangeOne)
	rangeTwo := scores[1:]
	fmt.Println(rangeTwo)
	rangeThree := scores[:3]
	fmt.Println(rangeThree)
	rangeThree = append(rangeThree, 100)
	fmt.Println(rangeThree)

}
