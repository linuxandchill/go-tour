package main

import "fmt"

type Person struct {
	name string
}

func (p Person) sayName() string {
	return p.name
}

func (p *Person) changeName() {
	p.name = "John"
}

func main() {
	p := Person{"Bill"}
	fmt.Println(p.name)
	fmt.Println(p.sayName())
	p.changeName()
	fmt.Println(p.sayName())
}
