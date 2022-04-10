package main

import "fmt"

func main() {
	fmt.Println(calculate("1+2*5/3+6/4*2"))
}

type stack struct {
	cap int
	s   []string
}

func NewStack(size int) *stack {
	return &stack{
		cap: size,
		s:   make([]string, 0, size),
	}
}

func (st *stack) Size() int {
	return len(st.s)
}

func (st *stack) Push(e string) {
	st.s = append(st.s, e)
}

func (st *stack) Pop() string {
	if st.Size() < 1 {
		panic("stack is empty!")
	}

	newSize := st.Size() - 1
	bf := st.s
	popped := bf[newSize]

	st.s = make([]string, 0, newSize)
	st.s = bf[0:newSize]
	return popped
}

func (st *stack) Top() string {
	if st.Size() < 1 {
		panic("stack is empty!")
	}
	return st.s[st.Size()-1]
}

func calculate(s string) int {
	operators := NewStack(len(s))

	var prev2, prev string
	for _, r := range s {
		char := string(r)
		if isOperator(char) {

		} else {
			prev2 = prev
			prev = char
		}
	}
}

var priority = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func isOperator(char string) bool {
	_, ok := priority[char]
	return ok
}

func operate(a, b int, operator string) int {
	ret := 0
	switch operator {
	case "+":
		ret = a + b
	case "-":
		ret = a - b
	case "*":
		ret = a * b
	case "/":
		ret = a / b
	}

	return ret
}
