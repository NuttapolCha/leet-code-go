package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "[dsaasd][{}]"
	fmt.Printf("isValid = %v\n", isValid(input))
}

type stack struct {
	l []string
}

func (st *stack) push(s string) {
	st.l = append(st.l, s)
}

func (st *stack) pull() string {
	ret := st.l[len(st.l)-1]
	st.l = st.l[0 : len(st.l)-1]
	return ret
}

func (st *stack) length() int {
	return len(st.l)
}

func isValid(s string) bool {
	is := true
	paranthesesStack := &stack{}
	openningParan := map[string]bool{
		"(": true,
		"{": true,
		"[": true,
	}
	closingParan := map[string]bool{
		")": true,
		"}": true,
		"]": true,
	}
	paranPairs := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	charactors := strings.Split(s, "")

	for _, c := range charactors {
		if isOpenParan := openningParan[c]; isOpenParan {
			paranthesesStack.push(c)
			continue
		}

		// len stack must not be 0 in order to pull the element from
		if isCloseParan := closingParan[c]; isCloseParan {
			// the found paran must be the OPPOSITE pulled out from stack
			if paranthesesStack.length() > 0 {
				paran := paranthesesStack.pull()
				if paranPairs[paran] != c {
					return false
				}
			} else {
				return false
			}
		}
	}

	// after checking each charactors, the final stack should be empty
	return is && paranthesesStack.length() == 0
}
