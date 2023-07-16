// Sharing Closures
package main

import "fmt"

type Calculator struct {
	add func(int, int) int
}

func NewCalculator() *Calculator {
	c := &Calculator{}
	c.add = func(x, y int) int {
		return x + y
	}
	return c
}

func (c *Calculator) Add(x, y int) int {
	return c.add(x, y)
}

func mains() {
	calc := NewCalculator
	fmt.Print(calc)
}
