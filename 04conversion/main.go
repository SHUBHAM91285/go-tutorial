package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our app")
	fmt.Println("Please rate our pizza app between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	//newreader allows the computer to read instruction entered by user

	input, _ := reader.ReadString('\n')
	//read until the user enters next line

	fmt.Println("Thanks for rating ", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //here 64 means float 64
	//if we don't use trimspace along with 4, \n is also passed so mathematical operation is not possible

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
