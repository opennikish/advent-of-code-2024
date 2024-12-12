package main

import "testing"

func TestSolve(t *testing.T) {
	cases := []struct {
		path     string
		expected int
	}{
		{
			path:     "test.txt",
			expected: 80,
		},
		{
			path:     "test2.txt",
			expected: 436,
		},
		{
			path:     "test3.txt",
			expected: 1206,
		},
		{
			path:     "test4.txt",
			expected: 236,
		},
		{
			path:     "test5.txt",
			expected: 368,
		},
	}

	for _, c := range cases {
		res := solve(c.path)
		if res != c.expected {
			t.Errorf("expected %d, got: %d", c.expected, res)
		}
	}
}
