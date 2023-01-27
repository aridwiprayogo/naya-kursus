package main

import "strings"

func FindShort(s string) int {
	// your code
	split := strings.Split(s, " ")

	lenght := len(split[0])
	for i := 0; i < len(split)-1; i++ {
		if lenght >= len(split[i]) {
			lenght = len(split[i])
		}
	}
	return lenght
}

func main() {
	x := "bitcoin take over the world maybe who knows perhaps"

	println(FindShort(x))
}
