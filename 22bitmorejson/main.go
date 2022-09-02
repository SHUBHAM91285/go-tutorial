package main

import (
	"encoding/json"
	"fmt"
)

//`json:"coursename"` -->> whosoever is consuming the json file will get coursename printed in the place of Name i.e. alias of Name
//`json:"-"` -->> whosoever is consuming the json file will not get password printed i.e. hide the password
//`json:"tags,omitempty"` -->> if there is nil then that place is not shown in json file
//there should not be gap between tags,omitempty i.e. `json:"tags, omitempty"` is not allowed

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json in golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamo", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamo", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamo", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	//package this data as json

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") //indent json based on tab. If we add 1 in the middle string we will get 1 at every line in output
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

//json data coming is in the form of bytes

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	
	{
		"coursename": "ReactJS Bootcamo",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
    }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json data is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse) // to print interfaces(or struct data) we use %#v
	} else {
		fmt.Println("JSON was not valid")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is: %T\n", k, v, v)
	}

	//order is not maintained in the key value pair
}
