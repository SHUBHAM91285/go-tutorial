package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in go")
	//no inheritance in golang
	Shubham := User{"Shubham Nath Tiwari", "shubham@go.dev", true, 20}
	fmt.Printf("Name is %v and email is %v.\n", Shubham.Name, Shubham.Email) //output ---> Name is Shubham Nath Tiwari and email is shubham@go.dev.
	Shubham.GetStatus()
	Shubham.NewMail() //it doesn't change the original email
	fmt.Printf("Email is %v.\n", Shubham.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "Test@go.dev" //here copy of email is passed not the original address. So,we need to use pointer to change original email
	fmt.Println("Email of this user is:", u.Email)
}

//methods are functions working on struct
//in java methods are functions inside class
//in c++ methods are function inside struct or class
