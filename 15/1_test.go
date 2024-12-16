package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<
`

	s := strings.TrimLeft(input, "\n")

	grid, moves := parseInput([]byte(s))

	expected := strings.TrimSpace(`
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########
`)

	g := bytes.Join(grid, []byte{'\n'})
	if expected != string(g) {
		t.Errorf("expected:\n%s\ngot:\n%s\n", expected, string(g))
	}

	m := "<^^>>>vv<v>>v<<"
	if len(m) != len(string(moves)) {
		t.Errorf("moves len diff, expected: %d, actual: %d", len(m), len(string(moves)))
	}
	if m != string(moves) {
		t.Errorf("expected:\n<^^>>>vv<v>>v<<\ngot:\n%s\n", string(moves))
	}
}

func TestMoveBoxes(t *testing.T) {
	input := strings.TrimSpace(`
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<
`)
	grid, moves := parseInput([]byte(input))

	moveBoxes(grid, moves)

	expected := strings.TrimSpace(`
########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########`)

	g := bytes.Join(grid, []byte{'\n'})

	if expected != string(g) {
		t.Errorf("expected:\n%s\ngot:\n%s\n", expected, string(g))
	}
}

func TestSolve(t *testing.T) {
	input := strings.TrimSpace(`
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<
`)

	res := solve([]byte(input))
	if res != 2028 {
		t.Errorf("expected 2028, got: %d", res)
	}
}
