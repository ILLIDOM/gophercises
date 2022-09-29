package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
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

// SELECT Pattern - (select to use first channel which is ready)
func channelSelect(c1 chan string, c2 chan string, c3 chan int) {
	time.Sleep(1 * time.Millisecond)
	select {
	case v1 := <-c1:
		fmt.Printf("received %v from c1\n", v1)
	case v2 := <-c2:
		fmt.Printf("received %v from c2\n", v2)
	case c3 <- 23:
		fmt.Printf("sent %v to c3\n", 23)
	default:
		fmt.Printf("no one was ready to communicate\n")
	}
}

func writeToChannel(c chan string) {
	c <- "message"
}

func concurrentDownload(urls []string, result chan string) {
	for _, url := range urls {
		go func(url string) {
			response, err := http.Get(url)
			if err != nil {
				result <- fmt.Sprintf("Failed downloading URL: %s\n", url)
			}
			defer response.Body.Close()
			if response.StatusCode != 200 {
				result <- fmt.Sprintf("Failed downloading URL: %s\n. Non 200 return code!", url)
			}
			var data bytes.Buffer
			_, err = io.Copy(&data, response.Body)
			if err != nil {
				result <- fmt.Sprintf("Failed to copy response Body of URL: %s", url)
			}
			result <- string(data.Bytes()[:40]) // write first 40 bytes as string into channel
		}(url)
	}

}

type Result struct {
	message string
	Error   error
}

func concurrentDownloadCustomResult(urls []string, result chan Result) {
	for _, url := range urls {
		go func(url string) {
			res := Result{}
			response, err := http.Get(url)
			if err != nil {
				res.Error = err
				result <- res
				return
			}
			defer response.Body.Close()
			if response.StatusCode != 200 {
				res.Error = errors.New(fmt.Sprintf("Failed downloading URL: %s\n. Non 200 return code!", url))
				result <- res
				return
			}
			var data bytes.Buffer
			_, err = io.Copy(&data, response.Body)
			if err != nil {
				res.Error = err
				result <- res
			}
			res.message = string(data.Bytes()[:40])
			result <- res // write first 40 bytes as string into channel
		}(url)
	}
}

func main() {
	// // synchronization without return value - WaitGroup
	// notify("Service-1", "Service-2", "Service-3")

	// // synchronization with return values - channels
	// notifyWithReturnValue("Service-1", "Service-2", "Service-3")

	// // communicate between main and a go-routine
	// c := make(chan string)
	// go printStuff("boring!", c)
	// for i := 0; i < 5; i++ { // read 5 strings from channel
	// 	fmt.Printf("You say: %q\n", <-c) // reading from channel - waits for sender (blocks)
	// }
	// fmt.Println("Leaving")

	// SELECT PATTERN
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan int)

	go writeToChannel(c1)
	go writeToChannel(c2)

	channelSelect(c1, c2, c3)

	// Async Download Simple
	urls := []string{"https://www.google.com", "https://www.20min.ch", "https://www.blick.ch", "https://start.fedoraproject.org/", ""}
	// result := make(chan string)
	// concurrentDownload(urls, result)

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(<-result)
	// }

	// Async Download custom response
	result := make(chan Result)
	concurrentDownloadCustomResult(urls, result)
	for i := 0; i < len(urls); i++ {
		res := <-result
		if res.message == "" {
			fmt.Printf("Nothing downloaded due to Error: %v\n", res.Error)
		} else {
			fmt.Printf("Downloaded: %s\n", res.message)
		}
	}

}
