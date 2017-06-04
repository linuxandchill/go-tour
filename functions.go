package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func mult(x, y string) (string, string) {
	return x, y
}

func namedRet(num int) (ten, six int) {
	ten = num + 10
	six = num + 6
	return
}

func main() {
	//fmt.Println(add(1, 2))
	//fmt.Println(mult("Hola", "dude"))
	fmt.Println(namedRet(12))
}
