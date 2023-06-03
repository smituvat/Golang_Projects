package main

import (
	"fmt"
	"sync"
	"time"
)

func count(s string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, s)
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)   // number of goroutines used
	go func() { // anonymous function
		count("smita")
		wg.Done()
	}()
	wg.Wait()
}
