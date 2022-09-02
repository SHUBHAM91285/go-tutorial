package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer in go")

	//defer prints the statement at last and also works on LIFO structure

	defer fmt.Println("This is a new statement")
	fmt.Println("Hello World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
