package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	// reverse those words
	split := strings.Split(str, " ")
	var result []string

	for i := len(split) - 1; i >= 0; i-- {
		result = append(result, split[i])
	}
	return strings.Join(result, " ")
}

func main() {
	fmt.Println(ReverseWords("hello world!"))
	fmt.Println(ReverseWords("yoda doesn't speak like this"))
	fmt.Println(ReverseWords("foobar"))
	fmt.Println(ReverseWords("kata editor"))
	fmt.Println(ReverseWords("row row row your boat"))
}
