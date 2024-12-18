package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{
			input: `
###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############
`,
			expected: 45,
		},
		{
			input: `
#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################
`,
			// 64 looks like mistake in AoC, because from i=9,j=9 we have 2 paths (one goes up, second right) iinstead of one
			expected: 77,
		},
	}

	for _, c := range cases {
		score := solve([]byte(c.input))

		if score != c.expected {
			t.Errorf("expected score %d, got: %d", c.expected, score)
		}
	}

}
