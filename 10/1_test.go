package main

import "testing"

func TestSolve(t *testing.T) {
	type Test struct {
		path     string
		expected int
	}

	for _, test := range []Test{
		{
			path:     "test.txt",
			expected: 36,
		},
		{
			path:     "test2.txt",
			expected: 2,
		},
		{
			path:     "test3.txt",
			expected: 4,
		},
	} {
		res := solve(test.path)
		if res != test.expected {
			t.Errorf("expected %d, got: %d", test.expected, res)
		}
	}

}
