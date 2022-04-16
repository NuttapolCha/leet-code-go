package main

import "fmt"

func main() {
	q := Constructor()
	q.Empty()
	q.Push(1)
	q.Push(2)
	q.Peek()
	q.Push(5)
	q.Pop()
	q.Empty()
	q.Peek()
	q.Push(99)
	q.Pop()
	q.Pop()
	q.Pop()
	q.Empty()
	// q.Pop()
	fmt.Printf("final queue = %+v\n", q.elementData)
}

type MyQueue struct {
	elementData []int
}

// Constructor allocate queue data structure using simple array with 10 slots
func Constructor() MyQueue {
	return MyQueue{
		elementData: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	// this condition might not occurred in this problem
	// if len(q.elementData) == q.size {
	// 	newSize := q.size + 1
	// 	newElementData := make([]int, 0, newSize)
	// 	q.size = newSize
	// 	for _, e := range q.elementData {
	// 		newElementData = append(newElementData, e)
	// 	}
	// 	q.elementData = newElementData
	// }
	q.elementData = append(q.elementData, x)
	fmt.Printf("queue after pushed: %d = %+v\n", x, q.elementData)
}

func (q *MyQueue) Pop() int {
	ret := q.Peek()

	old := q.elementData
	q.elementData = old[1:]

	fmt.Printf("queue after popped: %d = %+v\n", ret, q.elementData)
	return ret
}

func (q *MyQueue) Peek() int {
	if len(q.elementData) == 0 {
		panic("cannot Peek because queue is empty")
	}

	fmt.Printf("Peeked: %d\n", q.elementData[0])
	return q.elementData[0]
}

func (q *MyQueue) Empty() bool {
	fmt.Printf("is empty: %v\n", len(q.elementData) == 0)
	return len(q.elementData) == 0
}
