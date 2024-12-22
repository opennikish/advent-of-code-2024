package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	input := []byte(strings.TrimLeft(`
###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
`, "\n"))
	expected := 44
	res := solve(input, 2)
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
