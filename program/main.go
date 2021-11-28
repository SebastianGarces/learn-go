package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	myBill := createBill()
	promptOptions(myBill)
}

func getUserInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	response, error := r.ReadString('\n')
	response = strings.TrimSpace(response)
	return response, error
}

func promptOptions(b bill) string {
	reader := bufio.NewReader(os.Stdin)
	option, _ := getUserInput("Choose an option: \n a - add item \n s - save bill \n t - add tip\n", reader)
	switch option {
	case "a":
		name, _ := getUserInput("Item name: ", reader)
		price, _ := getUserInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)
		fmt.Printf("Item added: %v - %v\n", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getUserInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}

		b.updateTip(t)
		fmt.Printf("Tip added: %v\n", t)
		promptOptions(b)
	case "s":
		formattedBill := b.format()
		b.saveBill()
		fmt.Println(formattedBill)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}

	return option
}
