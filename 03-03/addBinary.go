package main

func addBinary(a string, b string) string {

	if len(a) > len(b) {
		return addBinary(b, a)
	}

	lenght := len(b) - len(a)

	for i := 0; i < lenght; i++ {
		a = "0" + a
	}

	carry := "0"
	result := ""

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '1' && b[i] == '1' {
			if carry == "1" {
				result = "1" + result
			} else {
				result = "0" + result
				carry = "1"
			}
		} else if a[i] == '0' && b[i] == '0' {
			if carry == "1" {
				result = "1" + result
				carry = "0"
			} else {
				result = "0" + result
			}
		} else if a[i] != b[i] {
			if carry == "1" {
				result = "0" + result
			} else {
				result = "1" + result
			}
		}
	}

	if carry == "1" {
		result = "1" + result
	}
	return result
}

func main() {
	println(addBinary("11", "1"))
	println(addBinary("1010", "1011"))
}
