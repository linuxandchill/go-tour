package main

import "fmt"

var list = []int{1, 2, 3, 4, 5}

func loopy(list []int) string {
	for i, value := range list {
		fmt.Println(i, value*2)
	}

	return "Done"
}

func main() {
	fmt.Println(loopy(list))
}
