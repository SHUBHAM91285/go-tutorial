package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang - LearnCodeOnline.in")

	//channels can interact with different go routines

	//chan is channel declaration chanel of int data type
	// mych := make(chan int)
	mych := make(chan int, 2) //buffered --- 2 channel(fixed) if we insert more than 2 value in channel
	wg := &sync.WaitGroup{}

	//insert 5 to mych <- arrow is used with channel
	// mych <- 5

	//gives error as deadlock (fatal error: all goroutines are asleep - deadlock!)
	// fmt.Println(<-mych)

	wg.Add(2)

	//RECEIVE ONLY
	//something is going out of the channel
	go func(ch <-chan int, wg *sync.WaitGroup) {

		close(mych)
		val, isChanelOpen := <-mych

		fmt.Println(isChanelOpen) //if mych<-0 then it will return true if mych<-0,mych<-1,mych<-2 etc. is not present then it will return false
		fmt.Println(val)

		// fmt.Println(<-mych)
		// fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)

	//listening to a closed channel is allowed
	//when we listen to a closed channel the output we get is zero
	//channels are bidirectional we can send and receive the value via channel

	//SEND ONLY
	//something is inserted in channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// mych <- 5
		// mych <- 6
		mych <- 0
		close(mych)
		wg.Done()
	}(mych, wg)

	wg.Wait()
}
