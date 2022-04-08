package main

import (
	"fmt"
	"sort"
)

func main() {
	// klar := Constructor(3, []int{4, 5, 8, 2})
	klar := Constructor(1, []int{})
	fmt.Printf("linked node = %v\n", toArray(klar.nodes))
	fmt.Printf("kth = %v\n", klar.kth)
	fmt.Println(klar.Add(-3))
	fmt.Println(klar.Add(-2))
	fmt.Println(klar.Add(-4))
	fmt.Println(klar.Add(0))
	fmt.Println(klar.Add(4))
}

type SortIntByAsc []int

func (a SortIntByAsc) Len() int           { return len(a) }
func (a SortIntByAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortIntByAsc) Less(i, j int) bool { return a[i] > a[j] }

type linkedNode struct {
	val  int
	next *linkedNode
}

type KthLargest struct {
	kth   int
	nodes *linkedNode
}

func Constructor(k int, nums []int) KthLargest {
	sorted := SortIntByAsc(nums)
	sort.Sort(sorted)

	// TODO: must refactor this section
	if len(nums) == 0 {
		return KthLargest{kth: k}
	}

	ret := KthLargest{
		kth:   k,
		nodes: &linkedNode{},
	}

	node := ret.nodes // init first node
	for i := 0; i < len(sorted); i++ {
		node.val = sorted[i]
		if i+1 < len(sorted) {
			node.next = &linkedNode{}
			node = node.next
		}
	}

	return ret
}

func (this *KthLargest) Add(val int) int {
	// fmt.Printf("=> original this: %v\n", toArray(this.nodes))
	iteration := 1
	ret := 0

	if this.nodes == nil {
		this.nodes = &linkedNode{val: val}
		// fmt.Printf("=> this after added: %d: %v\n", val, toArray(this.nodes))
		return val
	}

	node := this.nodes

	var isNewNodeInserted, isRetSet bool
	for node != nil {
		if !isNewNodeInserted {
			switch {
			case val > node.val:
				insertedNode := &linkedNode{
					val:  val,
					next: node,
				}
				this.nodes = insertedNode
				node = insertedNode
				isNewNodeInserted = true
			case val <= node.val && (node.next == nil || val >= node.next.val):
				insertedNode := &linkedNode{
					val:  val,
					next: node.next,
				}
				node.next = insertedNode
				isNewNodeInserted = true
			}
		}

		if iteration == this.kth {
			ret = node.val
			isRetSet = true
		}

		if isRetSet && isNewNodeInserted {
			break
		}

		iteration++
		node = node.next
	}

	// fmt.Printf("=> this after added: %d: %v\n", val, toArray(this.nodes))
	return ret
}

func toArray(node *linkedNode) []int {
	ret := []int{}
	for node != nil {
		ret = append(ret, node.val)
		node = node.next
	}

	return ret
}
