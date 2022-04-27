package main

import (
	"fmt"
)

func main() {
	pairs := [][]int{
		{8, 6},
		{3, 4},
		{5, 2},
		{5, 5},
		{3, 5},
		{7, 10},
		{6, 0},
		{10, 0},
		{7, 1},
		{4, 8},
		{6, 2},
	}

	fmt.Println(smallestStringWithSwaps("vbjjxgdfnru", pairs))
}

type minStringHeap struct {
	a   []string
	set map[string]bool
}

func NewMinStringHeap() minStringHeap {
	return minStringHeap{
		a:   make([]string, 0),
		set: make(map[string]bool),
	}
}

func (h *minStringHeap) Display() {
	for i, s := range h.a {
		fmt.Printf("index: %d => '%s'\n", i, s)
	}
}

func (h *minStringHeap) EnQueue(s string) {
	h.a = append(h.a, s)
	h.set[s] = true
	h.fixUp(len(h.a) - 1)
}

func (h *minStringHeap) DeQueue() string {
	out := h.Peek()

	// move the right most node to the first
	delete(h.set, out)
	h.a[0] = h.a[len(h.a)-1]
	h.a = h.a[:len(h.a)-1]
	h.fixDown(0)

	return out
}

func (h *minStringHeap) IsEmpty() bool {
	return len(h.a) == 0
}

func (h *minStringHeap) Peek() string {
	if h.IsEmpty() {
		panic("queue is empty!")
	}

	return h.a[0]
}

func (h *minStringHeap) fixUp(k int) {
	parentIndex := (k - 1) / 2
	if parentIndex < 0 {
		return
	}

	if h.a[k] < h.a[parentIndex] {
		h.a[k], h.a[parentIndex] = h.a[parentIndex], h.a[k]
		h.fixUp(parentIndex)
	}
}

func (h *minStringHeap) fixDown(k int) {
	leftChildIndex := 2*k + 1
	rightChildIndex := 2*k + 2

	if leftChildIndex < len(h.a) && h.a[leftChildIndex] < h.a[k] {
		h.a[leftChildIndex], h.a[k] = h.a[k], h.a[leftChildIndex]
		h.fixDown(leftChildIndex)
	}

	if rightChildIndex < len(h.a) && h.a[rightChildIndex] < h.a[k] {
		h.a[rightChildIndex], h.a[k] = h.a[k], h.a[rightChildIndex]
		h.fixDown(rightChildIndex)
	}
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	h := NewMinStringHeap()
	set := make(map[string]bool) // will stop swapping if swap result already exist in set

	// initial from input
	h.EnQueue(s)
	set[s] = true

	doSwap(&h, s, pairs)

	// h.Display()
	return h.Peek()
}

func doSwap(h *minStringHeap, s string, pairs [][]int) {
	// fmt.Printf("original s = %s\n", s)

	for _, pair := range pairs {
		index1 := pair[0]
		index2 := pair[1]
		// ensure that index1 always less than index2
		if index2 < index1 {
			index1, index2 = index2, index1
		}

		// swap chars
		if index1 != index2 {
			s = s[0:index1] + string(s[index2]) + s[index1+1:index2] + string(s[index1]) + s[index2+1:]
		}

		// fmt.Printf("after swap index (%d,%d) s = %s\n", index1, index2, s)

		if !h.set[s] { // not exist in set
			h.EnQueue(s)
			doSwap(h, s, pairs)
		}
	}
}
