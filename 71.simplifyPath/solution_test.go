package solution

import (
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	type testCase struct {
		testName string
		input    string
		expected string
	}

	testCases := []testCase{
		{
			testName: "1",
			input:    "/home/user/nuttapol/../kaimook/leet-code-go/../../nuttapol//leet-code-go/71.main.go/././...",
			expected: "/home/user/nuttapol/leet-code-go/71.main.go/...",
		},
		{
			testName: "2",
			input:    "/home/user/nuttapol/../kaimook/leet-code-go/../../nuttapol//leet-code-go/71.main.go/././...///../",
			expected: "/home/user/nuttapol/leet-code-go/71.main.go",
		},
		{
			testName: "3",
			input:    "/home//foo/",
			expected: "/home/foo",
		},
		{
			testName: "4",
			input:    "/home/",
			expected: "/home",
		},
		{
			testName: "5",
			input:    "/../",
			expected: "/",
		},
	}

	for _, tc := range testCases {
		actual := simplifyPath(tc.input)
		if tc.expected != actual {
			t.Fatalf("test: %s failed, expected: %s but got: %s\n", tc.testName, tc.expected, actual)
		}
	}
}
