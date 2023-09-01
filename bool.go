package main

import "fmt"

func main() {
	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos:", index)
			break

		}
		fmt.Printf("the value at Pos %v is %v\n", index, value)
	}

}
