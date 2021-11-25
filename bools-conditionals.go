package main

import "fmt"

func main() {
	age := 31

	fmt.Println(age <= 50) // true
	fmt.Println(age >= 50) // false
	fmt.Println(age == 31) // true
	fmt.Println(age != 50) // true

	if age < 35 {
		fmt.Println("Age is less than 35")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is not less than 45")
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos: ", index)
			continue
		}

		if index == 2 {
			fmt.Println("breaking at pos: ", index)
			break
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
}
