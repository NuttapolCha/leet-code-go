package solution

import "testing"

func TestWordFilter(t *testing.T) {
	words := []string{"apple", "banana", "orange", "car", "anaconde"}
	w := Constructor(words)

	type testCase struct {
		prefix   string
		suffix   string
		expected int
	}

	testCases := []testCase{
		{
			prefix:   "b",
			suffix:   "a",
			expected: 1,
		},
		{
			prefix:   "a",
			suffix:   "e",
			expected: 4,
		},
		{
			prefix:   "o",
			suffix:   "e",
			expected: 2,
		},
		{
			prefix:   "c",
			suffix:   "r",
			expected: 3,
		},
		{
			prefix:   "a",
			suffix:   "z",
			expected: -1,
		},
		{
			prefix:   "ap",
			suffix:   "e",
			expected: 0,
		},
		{
			prefix:   "b",
			suffix:   "na",
			expected: 1,
		},
		{
			prefix:   "oran",
			suffix:   "ge",
			expected: 2,
		},
		{
			prefix:   "anac",
			suffix:   "de",
			expected: 4,
		},
		{
			prefix:   "apple",
			suffix:   "apple",
			expected: 0,
		},
		{
			prefix:   "apple",
			suffix:   "le",
			expected: 0,
		},
		{
			prefix:   "ap",
			suffix:   "pple",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		actual := w.F(tc.prefix, tc.suffix)
		if tc.expected != actual {
			t.Errorf("prefix: %s, suffix: %s failed; expected: %d but got %d", tc.prefix, tc.suffix, tc.expected, actual)
		}
	}
}
