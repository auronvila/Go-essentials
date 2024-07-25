package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Go")

	channel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// pushing a 5 in the channel, you cannot push a value without listening to the channel
	// channel <- 5
	//fmt.Println(<-channel)

	wg.Add(2)
	// Receive Only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-channel
		fmt.Println(val)
		fmt.Println(isChannelOpen)
		wg.Done()
	}(channel, wg)

	// Send Only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		channel <- 5
		close(channel)
		wg.Done()
	}(channel, wg)

	wg.Wait()
}
