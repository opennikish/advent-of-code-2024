package main

import "testing"

func TestSolve(t *testing.T) {
	cases := []struct {
		path     string
		blinks   int
		expected int
	}{
		{
			path:     "test.txt",
			blinks:   6,
			expected: 22,
		},
		{
			path:     "test.txt",
			blinks:   25,
			expected: 55312,
		},
	}

	for _, c := range cases {
		res := solve(c.path, c.blinks)
		if res != c.expected {
			t.Errorf("expected %d, got: %d", c.expected, res)
		}
	}
}
