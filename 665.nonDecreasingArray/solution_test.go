package solution

import (
	"testing"
)

func TestCheckPossibility(t *testing.T) {
	type testCase struct {
		input    []int
		expected bool
	}

	testCases := []testCase{
		{
			input:    []int{1, 1, 2, 2, 5, 6, 7},
			expected: true,
		},
		{
			input:    []int{1, 1, 2, 1, 5, 6, 7}, // we can fix 3rd 1
			expected: true,
		},
		{
			input:    []int{1, 1, 2, 3, 5, 6, 3}, // we can fix the last 3
			expected: true,
		},
		{
			input:    []int{1, 1, 2, 3, 5, 6, 3, 8, 9, 1}, // we have to fix 2nd 3 and last 1
			expected: false,
		},
		{
			input:    []int{5, 4, 3}, // we have to fix 4 and 3
			expected: false,
		},
		{
			input:    []int{5, 4, 6}, // we can fix 4
			expected: true,
		},
		{
			input:    []int{5, 4, 6, 3, 2}, // we have to fix 4, 3, 2
			expected: false,
		},
		{
			input:    []int{5},
			expected: true,
		},
		{
			input:    []int{4, 2, 3}, // we can fix 4 to <=2
			expected: true,
		},
		{
			input:    []int{1, 4, 2, 3}, // we can fix 4
			expected: true,
		},
		{
			input:    []int{3, 4, 2, 3}, // we have to fix 2 and last 3
			expected: false,
		},
		{
			input:    []int{1, 3, 5, 2, 4}, // we have to fix 2 and 4
			expected: false,
		},
	}

	for _, tc := range testCases {
		originalInput := make([]int, len(tc.input))
		copy(originalInput, tc.input)

		actual := checkPossibility(tc.input)
		if actual != tc.expected {
			t.Errorf("input: %+v failed, expected %v but got %v, nums after fixed: %+v", originalInput, tc.expected, actual, tc.input)
		}
	}
}
