package main

import "fmt"

func adder() func(int) int { // outer function
	sum := 0
	return func(x int) int { // inner function
		sum += x
		return sum
	}
}

func main() {
	a := adder()
	fmt.Println(a(1))
	fmt.Println(a(2))
	fmt.Println(a(3))
}
