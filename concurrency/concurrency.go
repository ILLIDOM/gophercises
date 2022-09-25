package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// WaitGroup is more clear and can be used if no data will be returned from a go-routing
// wg.Add() should always be outside of concurrent code
// wg.Done() always inside of concurrent code -> else everythin is marked as done before the go-routine has run

func notify(services ...string) {
	var wg sync.WaitGroup
	for _, service := range services {
		wg.Add(1) // add 1 to wg for each service
		go func(s string) {
			fmt.Printf("Notifying service: %s\n", s)
			time.Sleep(time.Duration(3 * time.Second))
			fmt.Printf("Finished notifying service %s\n", s)
			wg.Done()
		}(service)
	}
	wg.Wait() // Blocking untill wg is zero
	fmt.Println("All services notified!")
}

// when a return value is needed from the go routines use channels
func notifyWithReturnValue(services ...string) {
	res := make(chan string)
	count := 0

	for _, service := range services {
		count++
		go func(s string) {
			fmt.Printf("Notifying service %s\n", s)
			time.Sleep(time.Duration(3 * time.Second))
			res <- fmt.Sprintf("Finished notifying service %s", s)
		}(service)
	}

	for i := 0; i < count; i++ {
		fmt.Println(<-res) // reading from a channel is a blocking call!
	}

	fmt.Println("All services notfied!")
}

// communication between main and go-routine using channels
func printStuff(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s, %d", msg, i) // write into channel - waits for receiver (blocks)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	// synchronization without return value - WaitGroup
	// notify("Service-1", "Service-2", "Service-3")

	// synchronization with return values - channels
	// notifyWithReturnValue("Service-1", "Service-2", "Service-3")

	// communicate between main and a go-routine
	c := make(chan string)
	go printStuff("boring!", c)
	for i := 0; i < 5; i++ { // read 5 strings from channel
		fmt.Printf("You say: %q\n", <-c) // reading from channel - waits for sender (blocks)
	}
	fmt.Println("Leaving")
}
