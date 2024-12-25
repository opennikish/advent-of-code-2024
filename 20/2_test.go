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
	expected := 285
	res := solve(input, 50)
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
