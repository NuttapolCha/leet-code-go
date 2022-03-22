package main

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	type testCase struct {
		roman string
		sum   int
	}

	testCases := []testCase{
		{
			roman: "I",
			sum:   1,
		},
		{
			roman: "II",
			sum:   2,
		},
		{
			roman: "III",
			sum:   3,
		},
		{
			roman: "IV",
			sum:   4,
		},
		{
			roman: "V",
			sum:   5,
		},
		{
			roman: "VI",
			sum:   6,
		},
		{
			roman: "VII",
			sum:   7,
		},
		{
			roman: "VIII",
			sum:   8,
		},
		{
			roman: "IX",
			sum:   9,
		},
		{
			roman: "X",
			sum:   10,
		},
		{
			roman: "XI",
			sum:   11,
		},
		{
			roman: "XIII",
			sum:   13,
		},
		{
			roman: "XIV",
			sum:   14,
		},
		{
			roman: "XVII",
			sum:   17,
		},
		{
			roman: "CDXXXIV",
			sum:   434,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.roman, func(t *testing.T) {
			actual := romanToInt(tc.roman)
			if actual != tc.sum {
				t.Errorf("expected: %d, but got: %d", tc.sum, actual)
			}
		})
	}

}
