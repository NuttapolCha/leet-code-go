package main

import (
	"fmt"
	"math"
)

func main() {
	lists := []*ListNode{
		nil, nil, nil, nil,
		nil,
		// {}, {},
		// {
		// 	Val: 0,
		// 	Next: &ListNode{
		// 		Val:  5,
		// 		Next: nil,
		// 	},
		// },
		// {
		// 	Val:  1,
		// 	Next: nil,
		// },
		// // {}, {}, {},
		// {
		// 	Val:  2,
		// 	Next: nil,
		// },
	}

	fmt.Println(mergeKLists(lists))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("%+v", l.ToArray())
}

func (l *ListNode) ToArray() []int {
	ret := make([]int, 0)

	next := l
	for next != nil {
		ret = append(ret, next.Val)
		next = next.Next
	}

	return ret
}

func newListNode(a []int) *ListNode {
	node := &ListNode{}

	next := node
	for i := 0; i < len(a); i++ {
		if i < len(a)-1 {
			next.Next = &ListNode{}
		}
		next.Val = a[i]
		next = next.Next
	}

	return node
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := func() int {
		sum := 0
		for _, list := range lists {
			next := list
			for next != nil {
				sum++
				next = next.Next
			}
		}
		return sum
	}()
	if length == 0 {
		return nil
	}

	ret := &ListNode{}
	next := ret

	for count := 0; count < length; count++ {
		next.Val = func() int {
			min := math.MaxInt
			indexMin := 0
			for i, list := range lists {
				if list != nil && list.Val < min {
					min = list.Val
					indexMin = i
				}
			}
			lists[indexMin] = lists[indexMin].Next
			return min
		}()
		if count < length-1 {
			next.Next = &ListNode{}
		}
		next = next.Next
	}

	return ret
}
