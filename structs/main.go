package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	fmt.Println(Person{"John", 24})
}
