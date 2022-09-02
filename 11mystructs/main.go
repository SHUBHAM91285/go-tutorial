package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in go")
	//no inheritance in golang
	Shubham := User{"Shubham Nath Tiwari", "shubham@go.dev", true, 20}
	fmt.Println(Shubham)                                                   //output ----> {Shubham Nath Tiwari shubham@go.dev true 20}
	fmt.Printf("Shubham's details are : %+v\n", Shubham)                   //output ---> Shubham's details are : {Name:Shubham Nath Tiwari Email:shubham@go.dev Status:true Age:20}
	fmt.Printf("Name is %v and email is %v.", Shubham.Name, Shubham.Email) //output ---> Name is Shubham Nath Tiwari and email is shubham@go.dev.
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
