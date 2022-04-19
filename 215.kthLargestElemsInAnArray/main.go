package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	// nums := []int{99, 99}
	sorted := []int{11, 10, 8, 8, 7, 7, 7, 6, 6, 6, 5, 5, 5, 5, 4, 4, 3, 3, 3, 2, 2, 2, 2, 1, 1, 1, 1}
	k := 20

	// forCheck := []int{}
	// for k := 1; k <= 20; k++ {
	// 	forCheck = append(forCheck, findKthLargest(nums, k))
	// }
	fmt.Printf("sorted = %v\n", sorted)
	// fmt.Printf("forche = %v\n", forCheck)

	fmt.Printf("actual: %d, expected: %d\n", findKthLargest(nums, k), sorted[k-1])
}

type minHeap struct {
	a    []int
	size int
}

func NewPriorityQueue(size int) minHeap {
	return minHeap{
		a:    make([]int, 0, size),
		size: size,
	}
}

func (q *minHeap) IsEmpty() bool {
	return len(q.a) == 0
}

func (q *minHeap) IsFull() bool {
	return len(q.a) == q.size
}

func (q *minHeap) EnQueue(v int) {
	if q.IsFull() {
		panic("cannot enqueue because queue is full!")
	}

	q.a = append(q.a, v)
	q.fixUp(len(q.a) - 1)
	// fmt.Printf("a after enqueue: %d = %v\n", v, q.a)
}

func (q *minHeap) Peek() int {
	if q.IsEmpty() {
		panic("cannot peek or dequeue because queue is empty!")
	}
	return q.a[0]
}

func (q *minHeap) DeQueue() int {
	ret := q.Peek()
	// q.a = q.a[1:]
	// TODO: understanding why we must move the right most node to the root node
	q.a[0] = q.a[len(q.a)-1]
	q.a = q.a[0 : len(q.a)-1]
	q.fixDown(0)
	// fmt.Printf("a after dequeue: %d = %v\n", ret, q.a)
	return ret
}

func (q *minHeap) fixDown(k int) {
	indexLeftChild := 2*k + 1
	indexRightChild := 2*k + 2

	if indexLeftChild <= len(q.a)-1 && q.a[k] > q.a[indexLeftChild] {
		q.a[k], q.a[indexLeftChild] = q.a[indexLeftChild], q.a[k]
		q.fixDown(indexLeftChild)
	}

	if indexRightChild <= len(q.a)-1 && q.a[k] > q.a[indexRightChild] {
		q.a[k], q.a[indexRightChild] = q.a[indexRightChild], q.a[k]
		q.fixDown(indexRightChild)
	}
}

func (q *minHeap) fixUp(k int) {
	if k < 1 {
		return
	}

	parentIndex := (k - 1) / 2

	if q.a[k] < q.a[parentIndex] {
		q.a[k], q.a[parentIndex] = q.a[parentIndex], q.a[k]
		q.fixUp(parentIndex)
	}
}

func findKthLargest(nums []int, k int) int {
	q := NewPriorityQueue(k)

	for _, num := range nums {
		if !q.IsFull() {
			q.EnQueue(num)
		} else {
			// fmt.Printf("kth has been reached, checking if %d >= %d or not\n", num, q.Peek())
			if num > q.Peek() {
				q.DeQueue()
				q.EnQueue(num)
			}
		}
	}

	return q.Peek()
}
