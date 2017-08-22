package main

import "fmt"

func main() {
	scores := []int{1, 2, 3, 5, 6}

	//remove last item
	scores = scores[:len(scores)-1]

	fmt.Println(scores)
}
