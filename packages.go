package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "f"))
	fmt.Println(strings.Split(greeting, " "))

	// the original value is not mutated, instead it returns the mutated value
	fmt.Println("the original string value:", greeting)

	ages := []int{30, 20, 45, 51, 24, 65, 91, 92}
	sort.Ints(ages) // mutates original slice
	fmt.Println(ages)

	// since slice is sorted, SearchInts returns index 0 for value 16
	// 16 being lower than first slice element but index 8 for value 95
	// 95 being higher than last slice element
	index := sort.SearchInts(ages, 16)  // 0
	index2 := sort.SearchInts(ages, 95) // 8
	fmt.Println(index, index2)
}
