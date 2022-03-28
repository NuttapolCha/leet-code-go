package main

import (
	"fmt"
	"sort"
)

func main() {
	// input := []int{-1, 0, 1, 2, -1, -4}
	// -4 -1 -1 0 1 2
	input := func() []int {
		return make([]int, 3000, 3000)
	}()
	fmt.Println(threeSum(input))
}

type intSorter []int

func (a intSorter) Len() int           { return len(a) }
func (a intSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a intSorter) Less(i, j int) bool { return a[i] < a[j] }

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)

	sortedNums := intSorter(nums)
	sort.Sort(sortedNums)

	appended := make(map[string]bool)

	var num2, num3 int
	for i, num1 := range sortedNums {
		for j := i + 1; j < len(sortedNums); j++ {
			num2 = sortedNums[j]
			for k := j + 1; k < len(sortedNums); k++ {
				num3 = sortedNums[k]
				if num1+num2+num3 == 0 {
					if !alreadyExist(appended, num1, num2, num3) {
						ret = append(ret, []int{num1, num2, num3})
					}
				}
			}
		}
		fmt.Println("end of iteration")
		fmt.Printf("len appended = %d, appended = %+v\n", len(appended), appended)
		// time.Sleep(time.Second * 1)
	}
	return ret
}

func alreadyExist(appended map[string]bool, num1, num2, num3 int) bool {
	if appended == nil {
		return false
	}

	key := fmt.Sprintf("%d:%d:%d", num1, num2, num3)
	if !appended[key] {
		appended[key] = true
		return false
	}

	return true
}
