package main

import (
	"fmt"
	"sort"
)

func main() {
	mat := [][]int{
		{1, 1, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}
	fmt.Println(kWeakestRows(mat, 3))
}

type scoreWithIndex struct {
	score int
	index int
}

type sortScoreWithIndexByScore []scoreWithIndex

func (a sortScoreWithIndexByScore) Len() int {
	return len(a)
}

func (a sortScoreWithIndexByScore) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a sortScoreWithIndexByScore) Less(i, j int) bool {
	if a[i].score == a[j].score {
		return a[i].index < a[j].index
	}
	return a[i].score < a[j].score
}

func (a sortScoreWithIndexByScore) indexes(k int) []int {
	ret := make([]int, 0, k)
	for i, score := range a {
		if i == k {
			break
		}
		ret = append(ret, score.index)
	}
	return ret
}

func kWeakestRows(mat [][]int, k int) []int {
	scores := make([]scoreWithIndex, 0, len(mat))
	for row, people := range mat {
		scoreOfThisRow := 0
		for _, person := range people {
			if person == 1 {
				scoreOfThisRow++
			} else {
				// no more soldiers in this row
				break
			}
		}
		scores = append(scores, scoreWithIndex{
			score: scoreOfThisRow,
			index: row,
		})
	}

	sortedScore := sortScoreWithIndexByScore(scores)
	sort.Sort(sortedScore)
	fmt.Printf("sorted = %+v\n", sortedScore)
	return sortedScore.indexes(k)
}
