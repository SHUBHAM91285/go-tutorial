package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Get request in golang")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:1111/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("The status code is: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	// content, _ := ioutil.ReadAll(response.Body)

	// fmt.Println(string(content))

	//other way to convert bytecode into string
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("The byte count is: ", byteCount)
	fmt.Println(responseString.String())
	fmt.Println(responseString.Len())
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:1111/post"

	//fake json payload(a payload describes the contents of a package)
	requestBody := strings.NewReader(`
	{
		"coursename":"let's go with golang.",
		"price":0,
		"platform":"learncodeonline.in"
	}
	`)

	//here in post request we have added url, type of data i.e. json, and body of data

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:1111/postform"

	//form data
	data := url.Values{}
	data.Add("FirstName", "Shubham")
	data.Add("MiddleName", "Nath")
	data.Add("LastName", "Tiwari")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
