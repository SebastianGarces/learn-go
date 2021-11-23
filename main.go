package main

import "fmt"

func main() {
	// print to console
	fmt.Println("Hello World!")

	// STRINGS
	var name string = "Sebastian"
	var lastName = "Garces" // infers type string
	var middleName string   // declare var to empty string (not undefined)
	fmt.Println(name, lastName, middleName)

	name = "Gerardo"
	lastName = "Barahona"
	middleName = "S"
	fmt.Println(name, middleName, lastName)

	// short hand
	// - only first time when declaring variable not for reassignment
	// - can only be used inside func body
	familyName := "Garces"
	fmt.Println("Family name:", familyName)

	// INTEGERS
	var age int = 31
	var miles int = 1542
	pieces := 40
	fmt.Println(age, miles, pieces)

	// bits and memory
	// see numeric types: https://go.dev/ref/spec#Numeric_types
	// 	uint8      the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers

	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts

	// byte        alias for uint8
	// rune        alias for int32
	var numOne int8 = 127
	var numTwo int16 = -129
	var numThree uint8 = 255 //unsigned numeric type (positive numbers)
	fmt.Println(numOne, numTwo, numThree)

}
