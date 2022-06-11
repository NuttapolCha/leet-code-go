package three

func lengthOfLongestSubstring(s string) int {
	latestIndex := make(map[rune]int)

	longest := 0
	latestDupIndex := -1
	curr := 0

	for i, char := range s {
		curr++
		if prev, is := latestIndex[char]; is {
			// check the current longest sub-string
			// we minus by 1 because we've add 1 for every the beginning of any iteration
			if prev < latestDupIndex {
				if curr > longest {
					longest = curr
				}
			} else {
				if curr-1 > longest {
					longest = curr - 1
				}
			}

			// adjust new curr
			if prev < latestDupIndex {
				curr = 1
			} else {
				curr = i - prev
			}
			latestDupIndex = i
		}
		latestIndex[char] = i
	}

	// check again
	if curr > longest {
		longest = curr
	}

	return longest
}
