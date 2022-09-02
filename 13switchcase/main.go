package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case in golang")

	rand.Seed(time.Now().UnixNano())
	//In the first line of the main function, we set seed to initialize a pseudorandom number generator.
	//The default math/rand number generator is deterministic, so it will give the same output sequence for the same seed value.
	//You can check this by removing the first line of the main function and running the program a couple of times
	//- we always get the same “random” numbers.
	//It is because the algorithm produces new values by performing some operations on the previous value,
	//and when the initial value (the seed value) stays the same, we get the same output numbers.
	//To avoid this, we use current time - time.Now().UnixNano() as the seed.

	diceNumber := rand.Intn(6) + 1 //Intn(6) returns number from 0 to 5
	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and now you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
		fallthrough //it means if case 4 is executed execute case 5 also
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot and you can roll dice again")
	default:
		fmt.Println("What was that!")
	}

	//no need to add break statement in every case like we did in c
}
