package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps in golang")
	languages := make(map[string]string) //the string in bracket represent key and other string represents value type

	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	fmt.Println("List of all languages: ", languages)
	//output is ----> List of all languages:  map[JS:javascript PY:python RB:ruby]

	fmt.Println("JS shorts for", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	//loops in map
	for key, value := range languages {
		fmt.Printf("For key %v value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("For key v value is %v\n", value)
	}
}
