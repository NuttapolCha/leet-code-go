package solution

func checkPossibility(nums []int) bool {
	hasFixed := false
	// this problem guarantee that len(nums) >= 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if hasFixed {
				return false
			} else {
				switch {
				// change left to curr
				case i-2 >= 0 && nums[i-2] <= nums[i], i-2 < 0:
					nums[i-1] = nums[i]
				// or change curr to left
				default:
					nums[i] = nums[i-1]
				}

				hasFixed = true
			}
		}
	}

	return true
}
