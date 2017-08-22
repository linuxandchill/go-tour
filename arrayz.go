package main

import "fmt"

func main() {
	//slice of ints, len of 0, cap of 10
	nums := make([]int, 0, 10)
	nums = append(nums, 5)
	fmt.Println(nums)

	//fmt.Println(cap(nums))

}
