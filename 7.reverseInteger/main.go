package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	x := 1234456
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	if float64(x) > math.Pow(2, 31)-1 || float64(x) < -1*math.Pow(2, 31) {
		return 0
	}

	var s string
	if x < 0 {
		s = fmt.Sprintf("%d", x*-1)
	} else {
		s = fmt.Sprintf("%d", x)
	}

	chars := strings.Split(s, "")

	var ret string
	if x < 0 {
		ret = "-"
	} else {
		ret = ""
	}
	for i := len(chars) - 1; i >= 0; i-- {
		ret += chars[i]
	}

	retInt, err := strconv.Atoi(ret)
	if err != nil {
		panic(err)
	}
	return retInt
}
