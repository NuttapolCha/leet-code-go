package main

import "fmt"

func main() {
	// l1 := NewListNodes([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	// l2 := NewListNodes([]int{5, 6, 4})

	l1 := NewListNodes([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := NewListNodes([]int{9, 9, 9, 9})

	// l1 := NewListNodes([]int{2, 4, 3})
	// l2 := NewListNodes([]int{5, 6, 4})

	digits1, _ := toDigits(l1)
	digits2, _ := toDigits(l2)

	fmt.Printf("l1 = %+v\n", digits1)
	fmt.Printf("l2 = %+v\n", digits2)

	result, _ := toDigits(addTwoNumbers(l1, l2))
	fmt.Println(result)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNodes(digits []int) *ListNode {
	ret := new(ListNode)
	next := ret

	for i, d := range digits {
		next.Val = d
		if i != len(digits)-1 {
			next.Next = new(ListNode)
			next = next.Next
		}
	}

	return ret
}

func toDigits(l *ListNode) ([]int, int) {
	ret := make([]int, 0)
	for l != nil {
		ret = append(ret, l.Val)
		l = l.Next
	}

	return ret, len(ret)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var additionalForNextVal int

	ret := new(ListNode)
	next := ret
	i := 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1 = new(ListNode)
		}
		if l2 == nil {
			l2 = new(ListNode)
		}

		sum := l1.Val + l2.Val + additionalForNextVal
		next.Val = sum % 10

		// prepare for next iteration
		additionalForNextVal = sum / 10

		// fmt.Printf("iteration: %d, l1.Val = %d, l2.Val = %d, additionalForNextVal = %d, sum = %d\n", i+1, l1.Val, l2.Val, additionalForNextVal, sum)

		l1 = l1.Next
		l2 = l2.Next

		// no more next, no more allocation
		if l1 != nil || l2 != nil {
			next.Next = new(ListNode)
			next = next.Next
		}
		i++
	}

	if additionalForNextVal != 0 {
		next.Next = new(ListNode)
		next = next.Next
		next.Val += additionalForNextVal
	}

	return ret
}
