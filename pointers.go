package main

import "fmt"

func main() {
	name := "sebastian"
	updateName(name)
	fmt.Println("name memory address: ", &name, "value: ", name)

	namePointer := &name
	fmt.Println("namePointer memory address: ", &namePointer, "value: ", *namePointer)

	updateNameByPointer(namePointer)
	fmt.Println("name memory address: ", &name, "value: ", name)
}

func updateName(name string) {
	name = "gerardo"
}

func updateNameByPointer(namePointer *string) {
	*namePointer = "gerardo"
}
