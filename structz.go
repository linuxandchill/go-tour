package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func main() {
	//use & to get address of value
	gohan := &Saiyan{"Gohan", 6900}
	//Super(gohan)
	//SuperPointer(gohan)
	gohan.SuperPointer()
	//fmt.Println(gohan.Power)
}

//parameter expected is now an address (*)
//*Saiyan is a receiver of SuperPointer method
func (s *Saiyan) SuperPointer() {
	s.Power += 1000
	fmt.Println(s.Power)
}

func Super(s Saiyan) {
	s.Power += 1000
}
