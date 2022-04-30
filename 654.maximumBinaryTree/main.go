package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	result := constructMaximumBinaryTree(nums)
	bs, _ := json.MarshalIndent(result, "", "\t")
	fmt.Println(string(bs))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index, max := findMax(nums)

	tree := &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[0:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}

	return tree
}

func findMax(nums []int) (index int, max int) {
	for i, n := range nums {
		if n > max {
			max = n
			index = i
		}
	}
	return
}

// type BinaryMaxHeap struct {
// 	a []int
// }

// func NewBinaryMaxHeap(nums []int) BinaryMaxHeap {
// 	h := BinaryMaxHeap{
// 		a: make([]int, 0, len(nums)),
// 	}

// 	for _, n := range nums {
// 		h.EnQueue(n)
// 	}

// 	return h
// }

// func (h *BinaryMaxHeap) IsEmpty() bool {
// 	return len(h.a) == 0
// }

// func (h *BinaryMaxHeap) EnQueue(v int) {
// 	h.a = append(h.a, v)
// 	h.fixUp(len(h.a) - 1)
// }

// func (h *BinaryMaxHeap) Peek() int {
// 	return h.a[0]
// }

// func (h *BinaryMaxHeap) DeQueue() int {
// 	ret := h.Peek()

// 	// replace the root node with  the right most node
// 	h.a[0] = h.a[len(h.a)-1]

// 	// resize the slice by ignoring the last element
// 	h.a = h.a[:len(h.a)-1]
// 	h.fixDown(0)

// 	return ret
// }

// func (h *BinaryMaxHeap) fixUp(k int) {
// 	parentIndex := (k - 1) / 2
// 	if parentIndex < 0 {
// 		return
// 	}

// 	if h.a[k] < h.a[parentIndex] {
// 		h.a[k], h.a[parentIndex] = h.a[parentIndex], h.a[k]
// 		h.fixUp(parentIndex)
// 	}
// }

// func (h *BinaryMaxHeap) fixDown(k int) {
// 	leftChildIndex := 2*k + 1
// 	rightChildIndex := 2*k + 2

// 	if leftChildIndex < len(h.a) && h.a[k] < h.a[leftChildIndex] {
// 		h.a[k], h.a[leftChildIndex] = h.a[leftChildIndex], h.a[k]
// 		h.fixDown(leftChildIndex)
// 	}

// 	if rightChildIndex < len(h.a) && h.a[k] < h.a[rightChildIndex] {
// 		h.a[k], h.a[rightChildIndex] = h.a[rightChildIndex], h.a[k]
// 		h.fixDown(rightChildIndex)
// 	}
// }
