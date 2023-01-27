package main

func GetSum(a, b int) int {
	if a == b {
		return a
	} else if a < b {
		return a + GetSum(a+1, b)
	}
	return a + GetSum(a-1, b)
}

func main() {
	println(GetSum(0, 1))      // 1
	println(GetSum(1, 2))      // 3
	println(GetSum(5, -1))     // 14
	println(GetSum(505, 4))    // 127759
	println(GetSum(321, 123))  // 44178
	println(GetSum(0, -1))     // -1
	println(GetSum(-50, 0))    // -1275
	println(GetSum(-1, -5))    // -15
	println(GetSum(-5, -5))    // -5
	println(GetSum(-505, 4))   // -127755
	println(GetSum(-321, 123)) // -44055
	println(GetSum(0, 0))      // 0
	println(GetSum(-5, -1))    // -15
	println(GetSum(5, 1))      // 15
	println(GetSum(-17, -17))  // -17
	println(GetSum(17, 17))    // 17
}
