package solution

import (
	"reflect"
	"testing"
)

func TestMaximumUniqueSubarray(t *testing.T) {
	type testCase struct {
		input    []int
		expected int
	}

	testCases := []testCase{
		{
			input:    []int{4, 2, 4, 5, 6},
			expected: 17,
		},
		{
			input:    []int{5, 2, 1, 2, 5, 2, 1, 2, 5},
			expected: 8,
		},
	}

	for _, tc := range testCases {
		actual := maximumUniqueSubarray(tc.input)
		if tc.expected != actual {
			t.Errorf("input %v have failed, expected: %d but got %d", tc.input, tc.expected, actual)
		}
	}
}

func TestGetPrefixSum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 3, 6, 10, 15}

	actual := getPrefixSum(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v but got: %v", expected, actual)
	}
}
