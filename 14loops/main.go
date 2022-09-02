package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to loops in go")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	//there is no pre-increment operator
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("The index is %v and day is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("The day is %v\n", day)
	// }

	rougeValue := 1

	//like while loop
	for rougeValue < 10 {

		// if rougeValue == 5 {
		// 	break
		// }

		// if rougeValue == 5 {
		// 	continue
		// }
		//it will wait upto 4 and will not proceed ahead as it does not know where to proceed and continue

		// if rougeValue == 5 {
		// 	rougeValue++
		// 	continue
		// }
		//here it knows that we have incremented the value now we can continue the loop

		if rougeValue == 2 {
			goto lco
		}
		//as it encounters 2 it jumps to lco tag

		fmt.Println(rougeValue)
		rougeValue++
	}
	//label tag (:) is used to give tag
lco:
	fmt.Println("Goto learncodeonline.in")
}
