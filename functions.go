package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	sayGreeting("Sebastian")
	sayBye("Oliver")

	names := []string{"mario", "luigi", "yoshi", "peach"}
	cycleNames(names, sayGreeting)
	cycleNames(names, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Println(a1, a2)
	fmt.Printf("circle a1 area is %0.2f and cicle a2 area is %0.2f \n", a1, a2)

	firstInitial, secondInitial := getInitials("Sebastian Garces")
	firstInitial2, _ := getInitials("Sebastian") // _ cannot be used as value or type
	fmt.Println(firstInitial, secondInitial)
	fmt.Println(firstInitial2)
}

func sayGreeting(name string) {
	fmt.Printf("Hi there, %v \n", name)
}

func sayBye(name string) {
	fmt.Printf("Bye %v \n", name)
}

func cycleNames(n []string, callback func(string)) {
	for _, value := range n {
		callback(value)
	}
}

func circleArea(radio float64) float64 {
	return math.Pi * radio * radio
}

func getInitials(name string) (string, string) {
	upperName := strings.ToUpper(name)
	names := strings.Split(upperName, " ")

	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}
