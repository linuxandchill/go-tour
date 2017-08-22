package main

import "fmt"

func main() {
	race := "asian"
	value, name := power(race)
	fmt.Printf("%s, %d, %s\n", name, value, race)
}

func power(race string) (int, string) {
	return 6, "John"
}
