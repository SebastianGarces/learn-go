package main

import "fmt"

func main() {

	// Print - same line console printing unless \n is added
	fmt.Print("hello, ")
	fmt.Print("world \n")
	fmt.Print("new line \n")

	// Println - auto new line
	fmt.Println("Hello Sebastian")
	fmt.Println("Goodbye Sebastian")

	age := 31
	name := "Sebastian"
	fmt.Println("my age is", age, "and my name is", name)

	// Formated strings - no new line
	// %_ = format specifier
	// %v = variable in order as remaining arguments
	// %q = for quotes - variable must be string
	// %T = prints variable type
	// %f = floats - can specified number of decimals for example %0.3f prints 3 decimal points
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %v and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)

	// Sprintf - returns formatted string
	var str = fmt.Sprintf("my age is %v and my name is %q \n", age, name)
	fmt.Println("var str:", str)
}
