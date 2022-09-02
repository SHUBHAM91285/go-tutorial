package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	//var ptr *int
	//fmt.Println("The value of pointer is ", ptr)
	//by default pointer is initialized as <nil>

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("The value of actual pointer is ", ptr)
	fmt.Println("The value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("The new value is: ", myNumber)
}
