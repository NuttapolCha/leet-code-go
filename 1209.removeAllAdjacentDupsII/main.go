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

type pair struct {
	char  string
	count int
}

type SpecialStack struct {
	pairs []pair
}

func newSpecialStack() SpecialStack {
	return SpecialStack{
		pairs: make([]pair, 0),
	}
}

func (st *SpecialStack) Push(s string) {
	st.pairs = append(st.pairs, pair{
		char: s,
		count: func() int {
			if st.IsEmpty() {
				return 1
			}
			if st.Peek().char == s {
				return 1 + st.Peek().count
			}
			return 1
		}(),
	})
}

func (st *SpecialStack) Peek() pair {
	return st.pairs[len(st.pairs)-1]
}

func (st *SpecialStack) Pop(k int) pair {
	ret := st.Peek()
	st.pairs = st.pairs[:len(st.pairs)-k]
	return ret
}

func (st *SpecialStack) BottomUp() string {
	ret := ""
	for _, pair := range st.pairs {
		ret += pair.char
	}

	return ret
}

func (st *SpecialStack) IsEmpty() bool {
	return len(st.pairs) == 0
}

func (st *SpecialStack) IsKDups(k int) bool {
	if st.IsEmpty() {
		return false
	}

	return st.Peek().count >= k
}

func removeDuplicates(s string, k int) string {
	st := newSpecialStack()

	chars := strings.Split(s, "")
	for _, char := range chars {
		st.Push(char)
		// fmt.Printf("pushed: %s, now st.pairs = %+v\n", char, st.pairs)
		// remove dups until no dups
		for st.IsKDups(k) {
			st.Pop(k)
		}
	}

	return st.BottomUp()
}
