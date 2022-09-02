package main

import "fmt"

func main() {
	fmt.Println("Welcomme to functions in golang")
	greeter()
	// greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	// proRes := proAdder(2, 5, 6, 9, 3)
	// fmt.Println("Result of proAdder is:", proRes)

	proRes, myMessage := proAdder(2, 5, 6, 9, 3)
	fmt.Println("Result of proAdder is:", proRes)
	fmt.Println("Result of proMessage is:", myMessage)
}

//function inside a function is not allowed

// func greeterTwo() {
// 	fmt.Println("Another method")
// }

func adder(valOne int, valTwo int) int { //int outside the bracket is return type of a function
	return valOne + valTwo
}

func greeter() {
	fmt.Println("Namastey golang")
}

//if we don't know how much parameter will be passed to a function we use proAdder function to add

// func proAdder(values ...int) int {
// 	total := 0
// 	for _, val := range values {
// 		total += val
// 	}
// 	return total
// }

//we can also return two types of value

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi pro message"
}
