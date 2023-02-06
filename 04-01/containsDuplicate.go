package main

import "sort"

func containsDuplicate(nums []int) (result bool) {
	sort.Ints(nums)
	for i := 1; i <= len(nums)-1; i++ {
		println(nums[i-1])
		if nums[i-1] == nums[i] {
			result = true
		}
	}
	return
}

func main() {
	println(containsDuplicate(
		[]int{3, 1, 2, 4}))
}
