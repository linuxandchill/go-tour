package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) sayName() string {
	return p.name
}

func main() {
	p := Person{"bill", 32}
	fmt.Println(p.sayName())
}
