package main

import "fmt"

func main() {
	s := "()()())(())(())()"
	fmt.Println(longestValidParentheses(s))
}

type stack struct {
	a []int
}

func (s *stack) push(i int) {
	s.a = append(s.a, i)
}

func (s *stack) peek() int {
	return s.a[len(s.a)-1]
}

func (s *stack) pop() int {
	i := s.peek()
	s.a = s.a[:len(s.a)-1]
	return i
}

func (s *stack) isEmpty() bool {
	return len(s.a) == 0
}

func longestValidParentheses(s string) int {
	// craate a stack topping with -1 as an initial state
	st := stack{}
	st.push(-1)

	maxLen := 0
	for i, c := range s {
		switch c {
		case '(':
			st.push(i)
		case ')':
			if !st.isEmpty() {
				st.pop()
				if !st.isEmpty() {
					currLen := i - st.peek()
					if currLen > maxLen {
						maxLen = currLen
					}
				} else {
					st.push(i)
				}
			} else {
				st.push(i)
			}

		default:
			panic("invalid charactor" + string(c))
		}
	}

	return maxLen
}
