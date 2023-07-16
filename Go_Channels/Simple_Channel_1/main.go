package main

import (
	"fmt"
	"sync"
)

// 1,2,3,4

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		Add(1)
		wg.Done()
	}()
	go func() {
		Add(3)
		wg.Done()
	}()
	wg.Wait()

}
func Add(n int) {
	fmt.Println(n + 5)
}
