package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 5}

	printy(s)

}

func printy(sli []int) {
	fmt.Printf("length: %d capacity: %d slice %v \n", len(sli), cap(sli), sli)

	new_sli := append(sli, 6, 7, 8)

	fmt.Printf("length: %d capacity: %d slice %v \n", len(new_sli), cap(new_sli), new_sli)

}
