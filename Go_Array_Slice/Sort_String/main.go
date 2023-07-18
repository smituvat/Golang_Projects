package main

import (
	"fmt"
	"sort" //import fmt and sort package
)

// create main function to execute the program
func main() {
	unsorted_str := "cbacdcbc" //create unsorted string
	fmt.Println("The unsorted string is:", unsorted_str)
	chars := []rune(unsorted_str)
	sort.Slice(chars, func(i, j int) bool { //sort the string using the function
		return chars[i] < chars[j]
	})
	fmt.Println("The sorted string is:")
	fmt.Println(string(chars)) //print the string on the console
	mp := make(map[rune]bool)
	chars1 := []rune{}
	for _, v := range chars {
		mp[v] = true
	}
	for k := range mp {
		chars1 = append(chars1, k)
	}
	fmt.Printf(string(chars1))
}
