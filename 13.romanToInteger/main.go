package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(romanToInt("XCIV")) // 94
}

func romanToInt(s string) (sum int) {
	chars := strings.Split(s, "")
	prev := ""
	for _, char := range chars {
		v := determineValue(prev, char)
		sum += v
		prev = char
	}

	return
}

func determineValue(prev, char string) int {
	isolateValues := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	prevValue := isolateValues[prev]
	currValue := isolateValues[char]

	if sub := subtractAmount(prev, char); sub != 0 {
		return currValue - prevValue - sub
	}
	return currValue
}

func subtractAmount(prev, char string) int {
	switch char {
	case "V", "X":
		if prev == "I" {
			return 1
		}
	case "L", "C":
		if prev == "X" {
			return 10
		}
	case "D", "M":
		if prev == "C" {
			return 100
		}
	}
	return 0
}
