package main

import "fmt"

// variable needs to be declared at top package
// level in order to be used by showScore
var score = 95

func main() {
	/*NOTES:
	Although IDE shows variables as undefined,
	when running main.go with command `go run`,
	if greetings.go is included as a command
	argument, it will run without any issues
	*/

	// TODO: find out how to export a package/module
	sayHello("Sebastian")
	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}
