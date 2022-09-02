package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	//Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.
	//Package os provides a platform-independent interface to operating system functionality.
	fmt.Println("Enter the rating of our pizza")
	//commma ok or comma err(,_)
	//used to store error message in a variable
	input, _ := reader.ReadString('\n')
	//rating is stored in input and error message  if any is stored in _
	//for _ we can use err also or for input we can use _ or any variable name.
	//if we use err in place of _ then we need to use it at any point in program
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
}
