package main

import "fmt"

type Car struct {
}

func (c Car) SUV(model string) bool {
	if model == "Explorer" {
		return true
	}
	return false
}

type Ford struct {
	Car
	model string
}

func main() {
	explorer := Ford{Car{}, "Explorer"}

	fmt.Println("Is it an Explorer?", explorer.SUV("Explorer"))
}
