package main

import (
	"bufio"
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	newBill := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return newBill
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	billName, _ := getUserInput("Create a new bill name: ", reader)
	newBill := newBill(billName)
	fmt.Println("Created bill: ", newBill.name)
	return newBill
}

// Receiver functions
func (b *bill) format() string {
	fs := "BILL BREAKDOWN: \n\n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("\n%-25v ...$%0.2f \n\n", "tip:", b.tip)
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)
	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip

}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) saveBill() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}
