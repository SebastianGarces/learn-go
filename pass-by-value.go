/*
Variable types (not all)
Group A			Group B
Strings			Slices
Ints			Maps
Floats			Functions
Booleans
Arrays
Structs
*/

package main

import "fmt"

func main() {
	/* Note: In type A variables Go passes arguments by value
	and not by reference therefore, original variable is not mutated */
	name := "sebastian"
	updateName(name)
	fmt.Println(name)

	/* returning and then reassigning value would work */
	name = updateNameReturn(name)
	fmt.Println(name)

	/* Note: In type B variables Go passes arguments by reference "pointer"
	therefore, reasignment does mutate original variable */
	menu := map[string]float64{
		"pie":       0.99,
		"ice cream": 3.99,
	}
	fmt.Println(menu)
	menu["pie"] = 5.99
	fmt.Println(menu)
}

func updateName(name string) {
	name = "gerardo"
}

func updateNameReturn(name string) string {
	name = "gerardo"
	return name
}
