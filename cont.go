package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"
	fmt.Println(greeting)
	fmt.Println(strings.Contains(greeting, "helo"))
	fmt.Println(strings.ReplaceAll(greeting, "friends", "my ghees"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "e"))
	fmt.Println(strings.Split(greeting, " "))

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 25)
	fmt.Println(index)

	unsorted := []string{"elvo", "eddy", "angel"}
	sort.Strings(unsorted)
	fmt.Println(unsorted)
	fmt.Println(sort.SearchStrings(unsorted, "eddy"))
}
