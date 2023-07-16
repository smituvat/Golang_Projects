package main

import "fmt"

func main() {
	fmt.Print(fibonacci_2(5))
}

// Recursive
func fibonacci_1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci_1(n-1) + fibonacci_1(n-2)
}

// Avoid repetations so store result in cache

var fibCache = make(map[int]int)

func fibonacci_2(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	if val, ok := fibCache[n]; ok {
		return val
	}

	result := fibonacci_2(n-1) + fibonacci_2(n-2)
	fibCache[n] = result
	return result
}
