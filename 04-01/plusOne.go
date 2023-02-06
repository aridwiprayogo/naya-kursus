package main

func plusOne(digits []int) []int {

	for i := range digits {
		if i == len(digits)-1 {
			if digits[i] >= 9 {
				digits[i] = 0
				digits[i-1] = digits[i-1] + 1
			}
			digits[i] += 1
		}
	}
	return digits
}

func main() {
	x := []int{1, 2, 3}
	for _, v := range plusOne(x) {
		println(v)
	}
}
