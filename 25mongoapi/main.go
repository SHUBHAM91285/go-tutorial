package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/SHUBHAM91285/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on port 4000 ...")
}
