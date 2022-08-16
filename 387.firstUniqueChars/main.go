package main

import "fmt"

func main() {
	s := "loveleetcodv"
	fmt.Println(firstUniqChar(s))
}

type pair struct {
	idx  int
	char rune
}

type queue struct {
	a []pair
}

func (q *queue) isEmpty() bool {
	return len(q.a) == 0
}

func (q *queue) enqueue(x pair) {
	q.a = append(q.a, x)
}

func (q *queue) peek() pair {
	return q.a[0]
}

func (q *queue) dequeue() pair {
	ret := q.peek()

	if len(q.a) > 1 {
		q.a = q.a[1:]
	} else {
		q.a = make([]pair, 0)
	}
	return ret
}

func firstUniqChar(s string) int {
	m := make(map[rune]int, 0)

	q := queue{}

	for idx, char := range s {
		if _, ok := m[char]; !ok {
			m[char] = 0
			q.enqueue(pair{idx, char})
		}
		m[char]++
	}

	for m[q.peek().char] > 1 {
		q.dequeue()
		if q.isEmpty() {
			break
		}
	}

	first := -1
	if !q.isEmpty() {
		first = q.peek().idx
	}
	return first
}
