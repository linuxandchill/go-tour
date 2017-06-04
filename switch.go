package main

import "fmt"

func main() {
	t := 20
	switch {
	case t < 12:
		fmt.Println("less than 12")
	default:
		fmt.Println(" 12")
	}
}
