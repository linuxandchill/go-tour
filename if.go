package main

import "fmt"

func test(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func main() {
	fmt.Println(test(12, 4))
}
