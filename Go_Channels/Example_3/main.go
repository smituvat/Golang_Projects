package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Routine 1"
			time.Sleep(time.Millisecond * 50)
		}
	}()

	go func() {
		for {
			c2 <- "Routine 2"
			time.Sleep(time.Second * 50)
		}
	}()

	for {
		// This is slow bcoz even though go routine 1 is fast still it waits for go routine 2
		// fmt.Println(<-c1)
		// fmt.Println(<-c2)

		//In order to solve this we can use select statement
		select {
		case m1 := <-c1:
			fmt.Println(m1)
		case m2 := <-c2:
			fmt.Println(m2) // go routine 1 will run multiple time as its faster
		}
	}
}
