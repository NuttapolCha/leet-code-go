package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(calculate("3+6 / 2*5"))
}

type StringStack struct {
	cap int
	s   []string
}

func NewStringStack(size int) *StringStack {
	return &StringStack{
		cap: size,
		s:   make([]string, 0, size),
	}
}

func (s *StringStack) size() int {
	return len(s.s)
}

func (s *StringStack) isEmpty() bool {
	return s.size() == 0
}

func (s *StringStack) push(char string) {
	if s.size() == s.cap {
		panic("could not push more in stack")
	}
	s.s = append(s.s, char)
}

func (s *StringStack) peek() string {
	if s.isEmpty() {
		return ""
	}
	return s.s[s.size()-1]
}

func (s *StringStack) pop() string {
	newSize := s.size() - 1
	ret := s.peek()
	bf := s.s

	s.s = make([]string, 0)
	s.s = bf[0:newSize]
	return ret
}

func calculate(infix string) int {
	return evalPostfix(infixToPostfix(infix))
}

func infixToPostfix(infix string) (postfix string) {
	chars := strings.Split(infix, "")
	st := NewStringStack(len(chars))

	for _, char := range chars {
		switch {
		case strings.TrimSpace(char) == "":
			continue
		case isOperator(char):
			switch {
			case st.isEmpty(), isHigherPriorityOperator(char, st.peek()):
				st.push(char)
			default:
				postfix += st.pop()
				st.push(char)
			}
		default:
			postfix += char
		}
	}

	for !st.isEmpty() {
		postfix += st.pop()
	}

	return
}

func evalPostfix(postfix string) int {
	fmt.Printf("before eval postfix: %s\n", postfix)
	chars := strings.Split(postfix, "")
	var prev, prev2 string
	var nextPostfix []string

	for i, char := range chars {
		if isOperator(char) {
			fmt.Printf("\t[iteration: %d is operator: %s]\n", i, char)
			a, err := strconv.Atoi(prev2)
			b, err := strconv.Atoi(prev)
			if err != nil {
				panic(fmt.Sprintf("invalid postfix: %s, curr = %s, prev = %s, prev2 = %s", postfix, char, prev, prev2))
			}
			result := operate(a, b, char)
			nextPostfix = append(nextPostfix[0:i-2], result)
			prev2 = string(nextPostfix[i-3])
			prev = result
		} else {
			nextPostfix = append(nextPostfix, char)
			prev2 = prev
			prev = char
		}
		fmt.Printf("\tdone iteration: %d: nextPostfix = %s\n", i, nextPostfix)
	}

	fmt.Printf("after eval postfix: %s\n", nextPostfix)
	return 0
}

var priority map[string]int = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func isOperator(char string) bool {
	return strings.Contains("+-*/", char)
}

func isHigherPriorityOperator(first, second string) bool {
	return priority[first] > priority[second]
}

func operate(a, b int, operator string) string {
	fmt.Printf("operating: %d %s %d\n", a, operator, b)
	switch operator {
	case "+":
		return fmt.Sprintf("%d", a+b)
	case "-":
		return fmt.Sprintf("%d", a-b)
	case "*":
		return fmt.Sprintf("%d", a*b)
	case "/":
		return fmt.Sprintf("%d", a/b)
	default:
		panic("invalid operator" + operator)
	}
}
