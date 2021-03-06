package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewCountMap(t *testing.T) {
	type testCase struct {
		input            []int
		expected         CountMap
		expectedScoreMap CountMap
	}

	testCases := []testCase{
		{
			input:            []int{1, 1, 2, 2, 3, 4, 5},
			expected:         map[int]int{1: 2, 2: 2, 3: 1, 4: 1, 5: 1},
			expectedScoreMap: map[int]int{1: 2, 2: 4, 3: 3, 4: 4, 5: 5},
		},
		{
			input:            []int{1, 1, 1, 1, 1, 0, 5},
			expected:         map[int]int{1: 5, 0: 1, 5: 1},
			expectedScoreMap: map[int]int{1: 5, 0: 0, 5: 5},
		},
		{
			input:            []int{22, 4, 4, 10, 95, 22},
			expected:         map[int]int{22: 2, 4: 2, 10: 1, 95: 1},
			expectedScoreMap: map[int]int{22: 44, 4: 8, 10: 10, 95: 95},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case: %d", i+1), func(t *testing.T) {
			actual, actualScoreMap := newCountMap(tc.input)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("count map expected: %v, but got: %v", tc.expected, actual)
			}
			if !reflect.DeepEqual(tc.expectedScoreMap, actualScoreMap) {
				t.Errorf("score map expected: %v, but got: %v", tc.expectedScoreMap, actualScoreMap)
			}
		})
	}
}
