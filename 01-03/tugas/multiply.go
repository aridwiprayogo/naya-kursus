package main

import "fmt"

func getMinLenght(numbers []int, lenght int) []int {

	for i := 1; i < len(numbers); i++ {
		numbers[i] = numbers[i-1] * numbers[i]
		if numbers[i] > lenght {
			break
		}
	}
	return numbers
}

func main() {
	number := []int{1, 2, 3, 4, 5}

	fmt.Println(getMinLenght(number, 9))
}
