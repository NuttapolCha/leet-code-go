package solution

func maximumUniqueSubarray(nums []int) int {
	numSet := make(map[int]bool)
	prefixSum := getPrefixSum(nums)
	latestDup := 0
	max := 0

	i := 0
	j := 0
	for j < len(nums) {
		// moving j til we found the duplicate
		for j < len(nums) {
			if !numSet[nums[j]] {
				// update the maximum score
				previousSum := 0
				if i > 0 {
					previousSum = prefixSum[i-1]
				}
				if score := prefixSum[j] - previousSum; score > max {
					max = score
				}
				// add nums[j] to set
				numSet[nums[j]] = true
				j++
			} else {
				latestDup = nums[j]
				break
			}
		}

		// exit the program if no more sub-array to check
		if j == len(nums) {
			break
		}

		// moving i til the latestDup
		for i < j {
			delete(numSet, nums[i])
			i++
			if nums[i-1] == latestDup {
				break
			}
		}
	}

	return max
}

func getPrefixSum(nums []int) []int {
	ret := make([]int, len(nums))
	prev := 0
	for i, num := range nums {
		ret[i] = num + prev
		prev = ret[i]
	}
	return ret
}
