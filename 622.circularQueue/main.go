package main

import "fmt"

func main() {
	q := Constructor(6)
	q.EnQueue(6)
	q.Rear()
	q.Rear()
	q.DeQueue()
	q.EnQueue(5)
	q.Rear()
	q.DeQueue()
	q.Front()
	q.DeQueue()
	q.DeQueue()
	q.DeQueue()
}

type MyCircularQueue struct {
	cap         int
	elementData []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		cap:         k,
		elementData: make([]int, 0, k),
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		printIf("cannot enqueue %d because queue is full", value)
		return false
	}

	q.elementData = append(q.elementData, value)
	printIf("queue after enqueue %d = %+v", value, q.elementData)
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		printIf("cannot dequeue because queue is empty")
		return false
	}

	dequeue := q.elementData[0]
	newElementData := q.elementData[1:]
	q.elementData = newElementData

	printIf("queue after dequeue %d = %+v", dequeue, q.elementData)
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	printIf("front is %d", q.elementData[0])
	return q.elementData[0]
}

func (q *MyCircularQueue) Rear() int {
	rear := len(q.elementData) - 1
	if rear < 0 {
		return -1
	}
	printIf("rear is %d", q.elementData[rear])
	return q.elementData[rear]
}

func (q *MyCircularQueue) IsEmpty() bool {
	printIf("IsEmpty = %v", len(q.elementData) == 0)
	return len(q.elementData) == 0
}

func (q *MyCircularQueue) IsFull() bool {
	printIf("IsFull = %v", len(q.elementData) == q.cap)
	return len(q.elementData) == q.cap
}

var debugMode = true

func printIf(f string, args ...interface{}) {
	if debugMode {
		fmt.Printf(f+"\n", args...)
	}
}
