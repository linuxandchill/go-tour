package main

import "fmt"

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func main() {
	gohan := &Saiyan{
		Name:  "Gohan",
		Power: 5000,
		Father: &Saiyan{
			Name:   "Goku",
			Power:  10000,
			Father: nil,
		},
	}
	//fmt.Println()
}
