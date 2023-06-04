package main

import (
	"fmt"
	"time"
)

func count(s string, c chan string) {
	for i := 1; i <= 3; i++ {
		c <- s
		time.Sleep(time.Millisecond * 50)
	}
	close(c) // to prevent deadlock we can use close channel, only sender can close the channel
}

func main() {
	c := make(chan string)
	go count("smita", c)

	// for { // this will go into deadlock because after 5 times count func will not have any value to send to chan but main still expects value

	for msg := range c {
		fmt.Println(msg)
	}
	// msg, open := <-c
	// if !open {
	// 	break
	// }
	// fmt.Println(msg)

	// }
}
