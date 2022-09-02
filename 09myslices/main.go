package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in go")

	var fruitList = []string{"Tomato", "Apple", "Peach"} //for slices we dont need to declare the size
	//we need to initialize the slice either by putting values or by simply using {} bracket e.g var fruitList = []string{}
	fmt.Printf("The data type of fruitlist is: %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana") //to add elements to slices
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:]) //reduce the original slice starting from index 1
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 555 ----> this will give error

	fmt.Println(highScores)
	highScores = append(highScores, 555, 666, 777)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores)) //here it will return false
	sort.Ints(highScores)                       //sorts the slice in increasing order
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) //returns true

	//to remove value from slices based on index
	var courses = []string{"reactjs", "javascript", "python", "ruby", "swift"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) //to remove elements of index 2
	fmt.Println(courses)
}
