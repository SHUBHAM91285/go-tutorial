package main

import "fmt"

//outside any method we cannot use := operator
//for e.g we cannot use user := 300000 outside main function
//but we can use var user = 300000 or var user int = 300000

const LoginToken string = "dgvhafchaxvC"

//if we start any variable name with capital letter it means the variable is public.

func main() {
	var username string = "Shubham Nath Tiwari"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.45545164665165616
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and aliases
	var anotherVariable int //default value is always zero
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type lexer automatically assign data type but we cannot change its data type by assigning
	//variable of other data type

	var website = "google.com"
	fmt.Println(website)
	//var website = 3; not allowed

	//no var style
	numberOfUsers := 30000
	fmt.Println(numberOfUsers)
	numberOfUser := 30000.2 //if we use 300000.0 it will print 300000 but for 300000.2 it will print 300000.2
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken) //%T is a placeholder for type of variable
	//%s is placeholder for string value.
}
