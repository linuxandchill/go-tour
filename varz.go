package main

import "fmt"

func main() {
	//var power int = 9000
	//power := 9000
	//fmt.Println(power)
	//fmt.Printf("It's over %d\n", power)
	hero, house := "goku", "mansion"
	power := getPower()
	fmt.Printf("%d\n", power)
	fmt.Printf("%s's house is a %s\n", hero, house)
}

func getPower() int {
	return 6900
}
