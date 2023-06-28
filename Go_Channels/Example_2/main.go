package main

import (
	"fmt"
)

func main() {
	/*
		// Ends up in deadlock bcoz send will block code until there is something to recive
		c := make(chan string)
		c <- "Hi"
		msg := <-c
		fmt.Println(msg)
	*/

	// To prevent this we can use buffered chan

	c := make(chan string, 2)
	c <- "hi"
	c <- "smita"

	msg := <-c
	msg = <-c
	fmt.Println(msg)
}
