package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "yfttttfbbbbnnnnffbgffffgbbbbgssssgthyyyy"
	k := 4
	fmt.Println(removeDuplicates(s, k))
}

type SpecialStack struct {
	chars []string
}

func newSpecialStack() SpecialStack {
	return SpecialStack{
		chars: make([]string, 0),
	}
}

func (st *SpecialStack) Push(s string) {
	st.chars = append(st.chars, s)
}

func (st *SpecialStack) Peek() string {
	return st.chars[len(st.chars)-1]
}

func (st *SpecialStack) Pop(k int) string {
	ret := st.Peek()
	st.chars = st.chars[:len(st.chars)-k]
	return ret
}

func (st *SpecialStack) BottomUp() string {
	return strings.Join(st.chars, "")
}

func (st *SpecialStack) IsEmpty() bool {
	return len(st.chars) == 0
}

func (st *SpecialStack) IsKDups(k int) bool {
	is := false

	if k <= len(st.chars) {
		prev := st.Peek()
		for i := 2; i <= k; i++ { // skip 1 because we already check from Peek()
			s := st.chars[len(st.chars)-i]
			is = s == prev
			if !is { // if false, break and return immediatly
				return is
			}
		}
	}

	return is
}

func removeDuplicates(s string, k int) string {
	st := newSpecialStack()

	chars := strings.Split(s, "")
	for _, char := range chars {
		st.Push(char)

		// remove dups until no dups
		for st.IsKDups(k) {
			st.Pop(k)
		}
	}

	return st.BottomUp()
}
