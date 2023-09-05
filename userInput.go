package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("create a new bill name :")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("Create a new bill name:", reader)
	b := newBill(name)
	fmt.Println("The created bill -", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose an option(a -add item ,s -save the bill ,t -add tip ):", reader)

	switch opt {
	case "a":
		// Implement logic to add an item to the bill
		fmt.Println("Adding an item...")
		name, _ := getInput("item name:", reader)
		price, _ := getInput("item price:", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be a number  ")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("the added item is :", name, price)
		promptOptions(b)
	case "s":
		b.saveBill()
		// Implement logic to save the bill
		fmt.Println(" you saved the bill:", b.name)
	case "t":
		// Implement logic to add a tip to the bill
		fmt.Println("Adding a tip...")
		tip, _ := getInput("enter a tip($):", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("the tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added", tip)
		promptOptions(b)
	case "q":
		return // Quit the function
	default:
		fmt.Println("Invalid option. Try again.")
		promptOptions(b)
	}

}

func main() {
	myBill := createBill()
	promptOptions(myBill)

}
