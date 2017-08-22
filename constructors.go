package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := new(Saiyan)
	gohan := &Saiyan{
		Name:  "gohan",
		Power: 5000,
	}

	fmt.Println(goku.Power)
	fmt.Println(gohan.Power)
}
