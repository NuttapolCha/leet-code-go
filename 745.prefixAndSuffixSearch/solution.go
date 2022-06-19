package solution

type WordFilter struct {
	m map[string]int
}

func Constructor(words []string) WordFilter {
	m := make(map[string]int)

	for i, word := range words {
		for fixedIndex := 0; fixedIndex <= len(word)-1; fixedIndex++ {
			for right := len(word) - 1; right >= 0; right-- {
				prefix := word[:fixedIndex+1]
				suffix := word[right:]
				m[prefix+":"+suffix] = i
				// fmt.Printf("m[%s:%s] = %d\n", prefix, suffix, i)
			}
			// fmt.Println()
		}
	}

	return WordFilter{
		m: m,
	}
}

func (w *WordFilter) F(prefix string, suffix string) int {
	key := prefix + ":" + suffix
	index, ok := w.m[key]
	if !ok {
		return -1
	}
	return index
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
