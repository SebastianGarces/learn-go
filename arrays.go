package main

import "fmt"

func main() {
	// Arrays have fixed lengths and cannot be changed
	var ages [3]int = [3]int{1, 2, 3}
	fmt.Println(ages, len(ages))
	list := [3]string{"chocolate", "banana", "steak"}
	fmt.Println(list, len(list))
	list[1] = "strawberry"
	fmt.Println(list, len(list))

	// Slices (use arrays under the hood)
	scores := []int{100, 50, 60}
	fmt.Println(scores, len(scores))
	// appends to end of slice - returns a new slice but can use reassignment to change initial slice
	scores = append(scores, 200)
	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne := list[1:3]  // includes start but excludes end
	rangeTwo := list[2:]   // no end value means from start index to end of slice
	rangeThree := list[:2] // no start value means from start of array to end value
	fmt.Println(rangeOne, len(rangeOne))
	fmt.Println(rangeTwo, len(rangeTwo))
	fmt.Println(rangeThree, len(rangeThree))

	rangeOne = append(rangeOne, "cheese")
	fmt.Println(rangeOne, len(rangeOne))
}
