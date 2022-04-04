package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode(nums []int) *ListNode {
	ret := new(ListNode)
	next := ret

	for i := 0; i < len(nums); i++ {
		next.Val = nums[i]
		// next node is not null
		if i != len(nums)-1 {
			next.Next = new(ListNode)
		}
		next = next.Next
	}

	return ret
}

func toArray(l *ListNode) []int {
	ret := make([]int, 0)

	for l != nil {
		ret = append(ret, l.Val)
		l = l.Next
	}
	return ret
}

func swapNodes(head *ListNode, k int) *ListNode {
	arr := toArray(head)

	fmt.Printf("arr = %v\n", arr)

	arr[k-1], arr[len(arr)-k] = arr[len(arr)-k], arr[k-1]

	fmt.Printf("arr = %v\n", arr)
	return newListNode(arr)
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	head := newListNode(input)

	swapped := swapNodes(head, 2)
	fmt.Println(toArray(swapped))
}
