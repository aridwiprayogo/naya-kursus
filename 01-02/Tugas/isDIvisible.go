package main

import "fmt"

func IsDivisible(n, x, y int) bool {
	// your code here
	return n%x == 0 && n%y == 0
}

func main() {
	fmt.Println(IsDivisible(3, 3, 4))
	fmt.Println(IsDivisible(12, 3, 4))
	fmt.Println(IsDivisible(8, 3, 4))
	fmt.Println(IsDivisible(48, 3, 4))
}
