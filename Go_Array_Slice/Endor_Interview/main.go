package main

import (
	"fmt"
)

func main() {
	res := fourthBit(90998)
	fmt.Print(res)
}

func fourthBit(number int32) int32 {
	// Write your code here
	var bin []int32
	var ans int32
	for number != 0 {
		bin = append(bin, number%2)
		number = number / 2
	}
	if len(bin) == 0 {
		return 0
	} else {
		a := len(bin) - 4
		ans = bin[a]
		// for i := len(bin); i >= 4; i-- {
		// 	ans = bin[i]
		// }
	}
	return ans
}
