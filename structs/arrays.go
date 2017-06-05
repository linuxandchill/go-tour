package main

import "fmt"

func main() {
	//var a as array of 10 ints
	var a [10]int

	a[0] = 0
	a[1] = 1
	fmt.Println(a)

	arr := []string{"hello", "lala"}
	fmt.Println(arr)

	// slice of structs
	str := []struct {
		make  string
		model string
	}{
		{"honda", "civic"},
	}

	fmt.Println(str[0])

}
