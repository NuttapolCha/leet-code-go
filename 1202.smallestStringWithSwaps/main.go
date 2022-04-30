package main

import (
	"fmt"
	"sort"
)

func main() {
	pairs := [][]int{
		{8, 6},
		{3, 4},
		{5, 2},
		{5, 5},
		{3, 5},
		{7, 10},
		{6, 0},
		{10, 0},
		{7, 1},
		{4, 8},
		{6, 2},
	}

	fmt.Println(smallestStringWithSwaps("vbjjxgdfnru", pairs))
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	linkedIndicesList := getLinkedIndices(pairs)

	smallestStr := s
	for _, linkedIndices := range linkedIndicesList {
		buff := getCharBuffersFromSelectedIndices(s, linkedIndices)
		swapedStr := sortStringOnlySelectedIndices(smallestStr, buff)
		if swapedStr < smallestStr {
			smallestStr = swapedStr
		}
	}

	return smallestStr
}

func getLinkedIndices(pairs [][]int) [][]int {
	// ret := make([][]int, 0)

	// linkedElems := make(map[int][]int, 0)
	// for _, pair := range pairs {
	// }
	return nil
}

type charBuffer struct {
	char          byte
	originalIndex int
	newIndex      int
}

func getCharBuffersFromSelectedIndices(s string, indices []int) []charBuffer {
	buff := make([]charBuffer, 0)

	for _, index := range indices {
		buff = append(buff, charBuffer{
			char:          s[index],
			originalIndex: index,
		})
	}

	return buff
}

func sortStringOnlySelectedIndices(s string, buff []charBuffer) string {
	chars := []byte(s)
	sort.Slice(buff, func(i, j int) bool {
		if buff[i].char < buff[j].char {
			buff[i].newIndex = j
			buff[j].newIndex = i
			return true
		}
		return false
	})

	for _, each := range buff {
		chars[each.originalIndex] = chars[each.newIndex]
	}

	return string(chars)
}
