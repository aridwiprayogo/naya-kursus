package main

import "fmt"

func MakeNegative(x int) (negatif int) {
	if x > 0 {
		negatif -= x
	} else {
		negatif = x
	}
	return
}

func main() {
	fmt.Println(MakeNegative(42))
	fmt.Println(MakeNegative(-9))
	fmt.Println(MakeNegative(0))
	fmt.Println(MakeNegative(1))
	fmt.Println(MakeNegative(-1))
}
