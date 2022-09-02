package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition - LearnCodeOnline.in")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{} //usually used when there is write condition i.e adding something to memory

	var score = []int{0}

	//anonymous function - function without name
	//receiving wg
	wg.Add(3)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("First go routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock() //unlock the resources
		wg.Done()
	}(wg, mut) //this brackets tells that run the function immediately,call wg and mut

	//wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Second go routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	//wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Third go routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut) //return it to wg and mut variable (global variable)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Third go routine read")
		mut.RLock() //read lock operation is performed whenever we use lock operation
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)
}

//if the condition go run --race .   gives no extra message in terminal it means everything is running good
