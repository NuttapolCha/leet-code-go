package main

import "fmt"

func main() {
	input := []byte("input")
	fmt.Println(input)
	reverseString(input)
	fmt.Println(input)
}

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1

	for left < right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp

		left++
		right--
	}
}
