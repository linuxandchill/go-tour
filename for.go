package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 5; i++ {
		sum += i
	}

	fmt.Println(sum)

	new_sum := 1
	//basically a while loop
	for new_sum < 1000 {
		new_sum += new_sum

	}
	fmt.Println(new_sum)
}
