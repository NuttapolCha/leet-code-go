package main

import "fmt"

func main() {
	heights := []int{2, 4}
	// heights := []int{2, 1, 5, 6, 2, 3, 99, 0, 98, 97, 96}
	// heights := []int{99}
	fmt.Println(largestRectangleArea(heights))
}

type pair struct {
	startIndex int
	val        int
}

// Stack is the increasing monotonic stack
type Stack struct {
	a []pair
}

func NewStack() Stack {
	return Stack{
		a: make([]pair, 0),
	}
}

func (st *Stack) IsEmpty() bool {
	return len(st.a) == 0
}

func (st *Stack) Push(startIndex, val int) {
	st.a = append(st.a, pair{startIndex, val})
}

func (st *Stack) Peek() pair {
	return st.a[len(st.a)-1]
}

func (st *Stack) Pop() (startIndex, val int) {
	ret := st.Peek()
	st.a = st.a[:len(st.a)-1]
	return ret.startIndex, ret.val
}

func largestRectangleArea(heights []int) int {
	st := NewStack()

	maximumArea := 0
	for i, height := range heights {
		switch {
		case st.IsEmpty():
			st.Push(i, height)

		case !st.IsEmpty() && height >= st.Peek().val:
			st.Push(i, height)

		case !st.IsEmpty() && height < st.Peek().val:
			// keep popping the heights till the top is equal or less than current height
			adjustedIndex := i
			for !st.IsEmpty() && height < st.Peek().val {
				// calculate area along with popping
				startIndex, prevHeight := st.Pop()
				area := prevHeight * (i - startIndex)
				if area > maximumArea {
					maximumArea = area
				}
				adjustedIndex = startIndex
			}
			st.Push(adjustedIndex, height)
		}
	}

	// finalize by calculate the remaining heights in stack
	for !st.IsEmpty() {
		i := len(heights) // note that i is out of heights' range
		startIndex, prevHeight := st.Pop()
		area := prevHeight * (i - startIndex)
		if area > maximumArea {
			maximumArea = area
		}
	}

	return maximumArea
}
