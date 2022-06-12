package three

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	longest := 0
	i := 0
	j := 0

	// do the following as we moving j
	// 1) update the longest length
	// 2) will stop moving j if we found dupclicate chars in a window
	var latestDup byte
	for j < len(s) {
		for j < len(s) {
			if charSet[s[j]] {
				latestDup = s[j]
				break
			}
			if length := j - i + 1; length > longest {
				longest = length
			}
			charSet[s[j]] = true
			j++
		}

		if j == len(s) {
			break
		}

		for i < j {
			delete(charSet, s[i])
			i++
			if s[i-1] == latestDup {
				break
			}
		}
	}

	return longest
}
