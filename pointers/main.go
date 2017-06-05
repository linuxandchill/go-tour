package main

import "fmt"

func main() {
	x := 32

	p := &x         // point to x
	fmt.Println(p)  //prints mem address
	fmt.Println(*p) // prints value at x
	// *p points to p's underlying value
	*p = 60 // change val of x to 60

	fmt.Println(x)
}
