package main

import (
	"fmt"
)

func main() {
	digits := "234"
	combs := letterCombinations(digits)
	fmt.Printf("len = %d\n", len(combs))

	for _, comb := range combs {
		fmt.Println(comb)
	}
}

var letterMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	ret := make([]string, 0)
	if len(digits) == 0 {
		return ret
	}
	if len(digits) == 1 {
		if letters, ok := letterMap[digits]; ok {
			return letters
		}
		return ret
	}

	// case len(digits) > 1:
	firstDigits := letterCombinations(digits[0:1])
	// fmt.Printf("firstDigits = %v\n", firstDigits)
	nextDigits := letterCombinations(digits[1:])
	// fmt.Printf("nextDigits = %v\n", nextDigits)
	for _, firstDigit := range firstDigits {
		for _, nextDigit := range nextDigits {
			ret = append(ret, firstDigit+nextDigit)
		}
	}

	return ret
}
