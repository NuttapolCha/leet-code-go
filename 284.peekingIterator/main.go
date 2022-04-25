package main

import "fmt"

func main() {
	i := &Iterator{
		a:     []int{1, 2, 3},
		index: 0,
	}

	p := Constructor(i)

	k := 0
	for p.hasNext() {
		fmt.Printf("k = %d, next = %d\n", k, p.peek())
		k++
		// p.next()
	}
}

type Iterator struct {
	a     []int
	index int
}

// Returns true if the iteration has more elements.
func (i *Iterator) hasNext() bool {
	fmt.Printf("current index = %d, has next = %v\n", i.index, i.index < len(i.a)-1)
	return i.index < len(i.a)-1
}

// Returns the next element in the iteration.
func (i *Iterator) next() int {
	i.index++
	return i.a[i.index]
}

type PeekingIterator struct {
	i *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{i: iter}
}

func (p *PeekingIterator) hasNext() bool {
	return p.i.hasNext()
}

func (p *PeekingIterator) next() int {
	return p.i.next()
}

func (p *PeekingIterator) peek() int {
	backup := *p.i
	ret := backup.next()
	return ret
}
