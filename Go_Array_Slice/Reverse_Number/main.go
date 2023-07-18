package main

import "fmt"

func main() {
	fmt.Print(reversed(45))
}

func reversed(n int32) int32 {
	var newInt int32
	for n > 0 {
		remainder := n % 10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	return newInt
}
