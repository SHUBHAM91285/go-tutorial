package main

import (
	"fmt"
	"net/url"
)

// fictional url
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hdshjb2454bnibcc"

func main() {
	fmt.Println("Welcome to urls")
	fmt.Println(myurl) //output ----> https://lco.dev:3000/learn?coursename=reactjs&paymentid=hdshjb2454bnibcc

	//parsing----> extracting information
	result, _ := url.Parse(myurl)
	fmt.Println(result)        //output ----> https://lco.dev:3000/learn?coursename=reactjs&paymentid=hdshjb2454bnibcc
	fmt.Println(result.Scheme) //output ----> https
	fmt.Println(result.Host)   //output ----> lco.dev:3000
	fmt.Println(result.Path)   //output ----> /learn
	//Port is a method
	fmt.Println(result.Port())   //output ----> 3000
	fmt.Println(result.RawQuery) //output ----> coursename=reactjs&paymentid=hdshjb2454bnibcc

	//Query() method stores parameter information about url

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams) //output ----> url.Values {url.value means key value pair}

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	//construct url from its part
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=shubham",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("The constructed url is: ", anotherUrl)
}
