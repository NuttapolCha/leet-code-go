package main

import "fmt"

func main() {
	testCases := [][]int{
		// {1, 0, 1, 1, 1, 1, 1, 3, 2, 1, 0, -1}, // true
		// {1, 2, 3, 4},                          // false
		// {3, 1, 4, 2},                          // true
		// {3, 5, 0, 3, 4},                       // true
		{-2, 1, 2, -2, 1, 2}, // true
	}
	for _, tc := range testCases {
		fmt.Println(find132pattern(tc))
	}
}

type stack struct {
	a []int
}

func (st *stack) size() int {
	return len(st.a)
}

func (st *stack) isEmpty() bool {
	return st.size() == 0
}

func (st *stack) peek() int {
	return st.a[len(st.a)-1]
}

func (st *stack) second() int {
	return st.a[len(st.a)-2]
}

func (st *stack) pop() int {
	ret := st.peek()
	st.a = st.a[:len(st.a)-1]
	// fmt.Printf("popped: %d\n", ret)
	return ret
}

func (st *stack) push(x int) {
	st.a = append(st.a, x)
	// fmt.Printf("pushed: %d\n", x)
}

func find132pattern(nums []int) bool {
	stks := make([]stack, len(nums))

	for i, num := range nums {
		stks[i].push(num)
		fmt.Printf("pushed: %d into stack[%d], current = %+v\n", num, i, stks[i])

		for j := 0; j < len(stks); j++ {
			switch {
			case i == j:
				continue
			case stks[j].size() == 1:
				if num > stks[j].peek() {
					stks[j].push(num)
					fmt.Printf("pushed: %d into stack[%d], current = %+v\n", num, j, stks[j])
				}
			case stks[j].size() == 2:
				if num < stks[j].peek() && num > stks[j].second() {
					return true
				}
			}
		}
	}

	return false
}
