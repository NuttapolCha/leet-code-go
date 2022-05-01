package main

import "fmt"

func main() {
	s := "#"
	t := "abe#c##########"
	fmt.Println(backspaceCompare(s, t))
}

func backspaceCompare(s string, t string) bool {
	return convert(s) == convert(t)
}

// convert converts # to a backspacing
func convert(s string) string {
	chars := []byte(s)

	ret := ""
	for _, b := range chars {
		if string(b) == "#" {
			if len(ret) > 0 {
				ret = ret[0 : len(ret)-1]
			}
		} else {
			ret += string(b)
		}
	}

	fmt.Printf("original = %s, converted = %s\n", s, ret)
	return ret
}
