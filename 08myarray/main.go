package main

import "fmt"

func main() {
	fmt.Println("Welcome to golang array")

	var fruitList [4]string //it is compulsary to give actual size of array

	fruitList[0] = "Apple"
	fruitList[1] = "Guava" //for fruitList[2] we have not added anything so it will print blank space for it
	fruitList[3] = "Peach"

	fmt.Println("The fruit list is: ", fruitList)
	fmt.Println("The fruit list is: ", len(fruitList))

	var vegList = [5]string{"Tomato", "Beans", "Mushroom"}
	fmt.Println("The fruit list is: ", vegList)
	fmt.Println("The fruit list is: ", len(vegList))
}
