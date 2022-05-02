package main

import "fmt"

func main() {
	nums1 := []int{}
	// nums2 := []int{3, 4, 5, 6, 7, 8}
	nums2 := []int{1, 2}

	fmt.Printf("%f\n", findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// if length is ood, the median is located at center
	lengthIsOdd := (len(nums1)+len(nums2))%2 == 1
	center := (len(nums1) + len(nums2)) / 2

	// running pointers for nums1 and nums2
	var i, j int
	var curr, prev int

	count := 0
	for count <= center {
		prev = curr
		switch {
		case i < len(nums1) && j < len(nums2):
			switch {
			case nums1[i] <= nums2[j]:
				curr = nums1[i]
				i++
			case nums1[i] > nums2[j]:
				curr = nums2[j]
				j++
			default:
				fmt.Println("what what")
			}
		case i < len(nums1): // only j is out of range
			curr = nums1[i]
			i++
		case j < len(nums2): // only i is out of range
			curr = nums2[j]
			j++
		default:
			fmt.Println("what")
		}

		count++
	}

	fmt.Printf("curr = %d, prev = %d, count = %d\n", curr, prev, count)

	var median float64
	if lengthIsOdd {
		median = float64(curr)
	} else {
		median = (float64(curr) + float64(prev)) / 2
	}

	return median
}
