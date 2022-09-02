package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //we always have to follow same format

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC) // create our own date
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday")) // month/date/year
	//to create an executable file for linux in windows system use the below command
	//$Env:GOOS = "linux"; $Env:GOARCH = "amd64"; go build
	//$Env:GOOS = "windows"; $Env:GOARCH = "amd64"; go build(to change from linux to windows to build .exe extension file)

}
