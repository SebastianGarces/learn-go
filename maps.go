package main

import "fmt"

func main() {
	/* maps must have same key/value types and they
	must be specified in map signature */
	menu := map[string]float64{
		"soup":            4.99,
		"pie":             7.99,
		"salad":           6.99,
		"soda":            1.99,
		"roasted-chicken": 9.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["roasted-chicken"])

	// LOOPS
	// TODO: find out why it doesn't loop in the same order every time
	for k, v := range menu {
		fmt.Printf("%v: %v \n", k, v)
	}

	phonebook := map[int]string{
		1234567890: "mario",
		2345678901: "luigi",
		3456789012: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[2345678901])

	phonebook[2345678901] = "sebastian"
	fmt.Println(phonebook)
}
