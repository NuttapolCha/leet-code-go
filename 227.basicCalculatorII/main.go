package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(calculate("2-0+30-56+0"))
	// fmt.Println(calculate("1*2-3/4+5*6-7*8+9/10"))
	// fmt.Println(calculate("1+2*5+2/2"))
	fmt.Println(calculate("1+2*5/3+6/4*2"))
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

func infixToPostfix(infix string) (postfix []string) {
	chars := splitExpression(infix)
	st := NewStringStack(len(chars))

	for _, char := range chars {
		switch {
		case isOperator(char):
			switch {
			case st.isEmpty(), isHigherPriorityOperator(char, st.peek()):
				st.push(char)
			case isEqualPriorityOperator(char, st.peek()):
				for isEqualPriorityOperator(char, st.peek()) {
					postfix = append(postfix, st.pop())
				}
				st.push(char)
			default: // i.e. lower priority
				for !st.isEmpty() {
					postfix = append(postfix, st.pop())
				}
				st.push(char)
			}
		default:
			postfix = append(postfix, char)
		}
	}

	for !st.isEmpty() {
		postfix = append(postfix, st.pop())
	}

	return
}

func evalPostfix(postfix []string) int {
	// fmt.Printf("postfix before eval: %s\n", postfix)

	var prev, prev2 string
	for i, char := range postfix {
		if isOperator(char) {
			a, err := strconv.Atoi(prev2)
			b, err := strconv.Atoi(prev)
			if err != nil {
				panic(fmt.Sprintf("invalid postfix: %s char = %s, a = %s, b = %s, err: %v", postfix, char, prev2, prev, err))
			}
			result := operate(a, b, char)
			nextPostfix := append(postfix[0:i-2], result)
			nextPostfix = append(nextPostfix, postfix[i+1:]...)
			return evalPostfix(nextPostfix)
		}
		prev2 = prev
		prev = char
	}

	// fmt.Printf("postfix after eval: %s\n", postfix)
	digits := ""
	for _, digit := range postfix {
		digits += digit
	}
	evaluated, err := strconv.Atoi(digits)
	if err != nil {
		panic(err)
	}
	return evaluated
}

func splitExpression(expression string) []string {
	ret := make([]string, 0)
	chars := strings.Split(expression, "")

	operand := ""
	for _, char := range chars {
		if strings.TrimSpace(char) == "" {
			// ignore any spaces
			continue
		}
		if isOperator(char) {
			ret = append(ret, operand)
			ret = append(ret, char)
			operand = ""
		} else {
			operand += char
		}
	}

	if operand != "" {
		ret = append(ret, operand)
	}

	// fmt.Printf("splitted expression: %s\n", ret)
	return ret
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

func isEqualPriorityOperator(first, second string) bool {
	return priority[first] == priority[second]
}

func operate(a, b int, operator string) string {
	// fmt.Printf("operating: %d %s %d\n", a, operator, b)
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
