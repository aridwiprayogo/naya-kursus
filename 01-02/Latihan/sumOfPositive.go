package main

import "fmt"

func PositiveSum(numbers []int) int {
	var result int
	for _, number := range numbers {
		if number > 0 {
			result += number
		}
	}
	return result
}

func main() {
	number := []int{1, 2, -3, -4, 5}
	fmt.Println(PositiveSum(number))
}
