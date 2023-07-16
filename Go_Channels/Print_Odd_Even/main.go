package main

import "fmt"

func main() {
	// var intSlice = []int{91, 42, 23, 14, 15, 76, 87, 28, 19, 95}
	chOdd := make(chan int)
	chEven := make(chan int)

	go odd(chOdd)
	go even(chEven)

	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			chOdd <- i
		} else {
			chEven <- i
		}
	}

}

func odd(ch <-chan int) {
	for v := range ch {
		fmt.Println("ODD :", v)
	}
}

func even(ch <-chan int) {
	for v := range ch {
		fmt.Println("EVEN:", v)
	}
}
