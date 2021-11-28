package main

import "fmt"

// go run structs.go build.go
func main() {
	myBill := createBill("sebastian's bill")
	fmt.Println(myBill)
}
