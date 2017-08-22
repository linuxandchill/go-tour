package main

import "fmt"

func main() {
	mappy := make(map[string]int)
	mappy["John"] = 21

	age, exists := mappy["John"]

	people := len(mappy)

	fmt.Println(age, exists)
	fmt.Println(people)

	cars := map[string]string{
		"Honda": "Accord",
		"Lambo": "Gallardo",
	}

	fmt.Println(cars["Honda"])
	for key, value := range cars {
		fmt.Println(key, value)
	}
}
