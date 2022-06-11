package three

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	type testCase struct {
		input    string
		expected int
	}

	testCases := []testCase{
		{
			input:    "tmmzuxt",
			expected: 5,
		},
		{
			input:    "abba",
			expected: 2,
		},
		{
			input:    "cdd",
			expected: 2,
		},
		{
			input:    "abcabcbb",
			expected: 3,
		},
		{
			input:    "bbbbb",
			expected: 1,
		},
		{
			input:    "pwwkew",
			expected: 3,
		},
		{
			input:    "a",
			expected: 1,
		},
		{
			input:    "dvdf",
			expected: 3,
		},
		{
			input:    "eabced",
			expected: 5,
		},
	}

	for _, tc := range testCases {
		actual := lengthOfLongestSubstring(tc.input)
		if actual != tc.expected {
			t.Errorf("input '%s' failed, expected %d but got %d", tc.input, tc.expected, actual)
		}
	}
}
