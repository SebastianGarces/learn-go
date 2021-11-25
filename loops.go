package main

import "fmt"

func main() {

	x := 0
	for x < 5 {
		fmt.Println("value of x is: ", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is: ", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("position at index %v is %v \n", index, value)
	}

	// using _ for unwanted variables - notice no unused variable error
	for _, value := range names {
		fmt.Printf("value is %v \n", value)
		value = "some random value"
		// this does not mutate original slice as value is a local
		// copy within the loops execution context
	}

	fmt.Println(names)
}
