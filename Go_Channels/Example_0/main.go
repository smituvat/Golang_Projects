package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	// Print statement keeps waiting untill msg reives value so its used in snchronizing go routines
	fmt.Println(msg)
}

// sending and receiving operations on channels block until both the sender and receiver are ready
// Link :  https://medium.com/@isuruvihan/exploring-channels-in-go-connecting-goroutines-2f61c39015b8

// Unlike unbuffered channels, which require both the sender and receiver to be ready simultaneously, buffered channels provide temporary
// storage for values until a receiver is available.
// This buffering capability allows the sender to continue sending values even if the receiver is not ready.

//example

func mains() {
	message := make(chan string, 1)
	done := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		message <- "Hello, buffered channel!"
		done <- true
	}()

	select {
	case msg := <-message:
		fmt.Println(msg)
	case <-done:
		fmt.Println("Task completed!")
	}
}

// ch := make(chan<- int) // Creating a send-only channel
// ch := make(<-chan int) // Creating a receive-only channel
