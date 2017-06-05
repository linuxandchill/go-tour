package main

import "fmt"

var list = []int{1, 2, 3, 4, 5, 6, 7}

func main() {
	for i, v := range list {
		fmt.Printf("2**%d = %d \n", i, v)
	}
}
