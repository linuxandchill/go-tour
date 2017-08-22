package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, my name is %s\n", p.Name)
}

func main() {
	john := &Person{
		Name: "John",
	}

	john.Introduce()
}
