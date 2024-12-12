package main

import "testing"

func TestSolve(t *testing.T) {
	cases := []struct {
		path     string
		expected int
	}{
		{
			path:     "test.txt",
			expected: 140,
		},
		{
			path:     "test2.txt",
			expected: 772,
		},
		{
			path:     "test3.txt",
			expected: 1930,
		},
	}

	for _, c := range cases {
		res := solve(c.path)
		if res != c.expected {
			t.Errorf("expected %d, got: %d", c.expected, res)
		}
	}
}
