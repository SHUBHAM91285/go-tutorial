package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

func main() {
	// go greeter("Hello")
	//go key word is used for goroutines
	//we are firing the thread
	// greeter("World")

	website := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range website {
		wg.Add(1) //keeps on adding the number of threads being fired
		getStatusCode(web)
	}

	//always written at the end of main function
	wg.Wait() //doesn't allow the main signal to end until the done signal is send
	fmt.Println(signals)
}

// func greeter(s string) {
// for i := 0; i < 6; i++ {
//here we are receiving the thread
// time.Sleep(5 * time.Second)
// 	time.Sleep(5 * time.Millisecond)
// 	fmt.Println(s)
// }
// }

/* to understand more about goroutines use
the website: https://golangbot.com/goroutines/
*/

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	//endpoint receives the website
	if err != nil {
		fmt.Println("OOPs! in status code")
	} else {
		mut.Lock() //doesn,t allow other process to use memory unless it completes its process
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d is the status code for %s\n", res.StatusCode, endpoint)
	}
}
