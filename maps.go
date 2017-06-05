package main

import "fmt"

type Car struct {
	model string
}

var c map[string]Car

func main() {
	c = make(map[string]Car)
	c["Honda"] = Car{
		"Civic",
	}

	fmt.Println(c["Honda"])
}
