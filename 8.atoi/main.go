package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("-42"))
	fmt.Println(myAtoi("  99"))
	fmt.Println(myAtoi("  9o9"))
	fmt.Println(myAtoi("-4   2"))
	fmt.Println(myAtoi("    34dfs42"))
	fmt.Println(myAtoi("934"))
	fmt.Println(myAtoi("  wo rd and 2"))
	fmt.Println(myAtoi("-42 with words"))
	fmt.Println(myAtoi("+42  dd   234 d"))
}

var intMap = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func myAtoi(s string) int {
	var isReading bool
	var isNegative bool
	var toBeRead string

	for _, b := range s {
		c := string(b)
		_, isDigit := intMap[c]
		if isReading {
			if isDigit {
				toBeRead += c
			}
			if c == "." {
				break
			}
		} else {
			if !isDigit && c != " " && c != "+" && c != "-" {
				return 0
			}
			if isDigit || c == "+" || c == "-" {
				isReading = true
				if isDigit {
					toBeRead += c
				}
				if c == "-" {
					isNegative = true
				}
			}
		}
	}

	var ret, pow float64
	for i := len(toBeRead) - 1; i >= 0; i-- {
		ret += float64(intMap[string(toBeRead[i])]) * math.Pow(10, pow)
		pow += 1
	}

	if ret > 2147483648 {
		ret = 2147483648
	}

	if isNegative {
		ret = ret * -1
	}

	// fmt.Printf("toBeRead = %s, isNegative = %v\n", toBeRead, isNegative)
	return int(ret)
}
