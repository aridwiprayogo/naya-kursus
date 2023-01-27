package main

import "fmt"

func Hero(bullets, dragons int) bool {
	// your code
	return bullets >= (dragons * 2)
}

func main() {
	fmt.Println(Hero(10, 5) == true)
	fmt.Println(Hero(7, 4) == false)
	fmt.Println(Hero(4, 5) == false)
	fmt.Println(Hero(100, 40) == true)
	fmt.Println(Hero(1500, 751) == false)
	fmt.Println(Hero(0, 1) == false)
}
