package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(scoreOfParentheses("()((())()())")) // 9
}

type stack struct {
	l []int
}

func (s *stack) push(i int) {
	s.l = append(s.l, i)
}

func (s *stack) pop() int {
	lastElem := s.l[len(s.l)-1]
	s.l = s.l[0 : len(s.l)-1]
	return lastElem
}

func scoreOfParentheses(s string) int {
	chars := strings.Split(s, "")

	scoreStack := &stack{}
	score := 0

	var prevIsOpen bool
	for _, char := range chars {
		switch {
		case isOpen(char):
			scoreStack.push(score)
			score = 0
			prevIsOpen = true
		case isClosed(char):
			score = scoreStack.pop() + func() int {
				if prevIsOpen {
					prevIsOpen = false
					return 1
				}
				return 2 * score
			}()
		}
	}

	return score
}

func isOpen(c string) bool {
	return c == "("
}

func isClosed(c string) bool {
	return c == ")"
}
