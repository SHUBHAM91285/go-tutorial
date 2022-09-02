package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go")
	content := "This needs to go in file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt") //create file in working directory

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content) //it gives length of string and also copy the content in file
	checkNilErr(err)

	fmt.Println("Length  is:", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

//whenever we read the file it's always read in bytes

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Byte data inside file is\n", databyte)
	fmt.Println("Text data inside file is\n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err) //stops the execution of program and displays error message
	}
}
