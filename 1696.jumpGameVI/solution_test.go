package solution

import "testing"

func TestMaxResult(t *testing.T) {
	type testCase struct {
		nums     []int
		k        int
		expected int
	}

	testCases := []testCase{
		{
			nums:     []int{1, -1, -2, 4, -7, 3},
			k:        2,
			expected: 7,
		},
		{
			nums:     []int{10, -5, -2, 4, 0, 3},
			k:        3,
			expected: 17,
		},
		{
			nums:     []int{1, -5, -20, 4, -1, 3, -6, -3},
			k:        2,
			expected: 0,
		},
	}

	for i, tc := range testCases {
		actual := maxResult(tc.nums, tc.k)
		if actual != tc.expected {
			t.Errorf("error on test case no: %d, expected: %d but got: %d", i+1, tc.expected, actual)
		}
	}
}
