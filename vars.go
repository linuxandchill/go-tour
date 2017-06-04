package main

import "fmt"

var y int = 30

func changeY(y *int) {
	*y += 20
}

func main() {
	x := 12

	fmt.Println(x)

	fmt.Println(y)
	//y += 20
	changeY(&y)
	fmt.Println(y)
}
