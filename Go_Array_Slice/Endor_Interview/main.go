package main

import "fmt"

func main() {

	res := make([]string, 2)
	res = append(res, "A")
	res = append(res, "A")
	fmt.Println(res)
}
