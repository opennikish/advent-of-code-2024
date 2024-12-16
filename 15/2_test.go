package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `
##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<<vv>^<v^>v>^v
v^v>v<>v^v<v<^
`

	s := strings.TrimLeft(input, "\n")

	grid, moves := parseInput([]byte(s))

	expected := strings.TrimSpace(`
####################
##....[]....[]..[]##
##............[]..##
##..[][]....[]..[]##
##....[]@.....[]..##
##[]##....[]......##
##[]....[]....[]..##
##..[][]..[]..[][]##
##........[]......##
####################
`)

	g := bytes.Join(grid, []byte{'\n'})
	if expected != string(g) {
		t.Errorf("expected:\n%s\ngot:\n%s\n", expected, string(g))
	}

	m := "<<vv>^<v^>v>^vv^v>v<>v^v<v<^"
	if len(m) != len(string(moves)) {
		t.Errorf("moves len diff, expected: %d, actual: %d", len(m), len(string(moves)))
	}
	if m != string(moves) {
		t.Errorf("expected:\n%s\ngot:\n%s\n", m, string(moves))
	}
}

func TestMoveBoxes(t *testing.T) {
	input := strings.TrimSpace(`
##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
`)
	grid, moves := parseInput([]byte(input))

	moveBoxes(grid, moves)

	expected := strings.TrimSpace(`
####################
##[].......[].[][]##
##[]...........[].##
##[]........[][][]##
##[]......[]....[]##
##..##......[]....##
##..[]............##
##..@......[].[][]##
##......[][]..[]..##
####################`)

	g := bytes.Join(grid, []byte{'\n'})

	if expected != string(g) {
		t.Errorf("expected:\n%s\ngot:\n%s\n", expected, string(g))
	}

	sum := calcPointsSum(grid)
	if sum != 9021 {
		t.Errorf("expected sum 9021, got: %d", sum)
	}
}

func TestMove(t *testing.T) {
	cases := []struct {
		grid     string
		moves    string
		expected string
	}{
		{
			grid: `
##########
#........#
#.[].....#
#.[].....#
#..@.....#
##########
`,
			moves: "^",
			expected: `
##########
#.[].....#
#.[].....#
#..@.....#
#........#
##########
`,
		},
		{
			grid: `
##########
#..#.....#
#.[].....#
#.[].....#
#..@.....#
##########
`,
			moves: "^",
			expected: `
##########
#..#.....#
#.[].....#
#.[].....#
#..@.....#
##########
`,
		},
		{
			grid: `
##########
#........#
#..[]....#
#.[][]...#
#.@......#
##########
`,
			moves: "^",
			expected: `
##########
#..[]....#
#.[].....#
#.@.[]...#
#........#
##########
`,
		},
		{
			grid: `
##########
#........#
#..[]....#
#.[][]...#
#..[]....#
#..@.....#
##########
				`,
			moves: "^",
			expected: `
##########
#..[]....#
#.[][]...#
#..[]....#
#..@.....#
#........#
##########
`,
		},
		{
			grid: `
##########
#........#
#..[]#...#
#.[][]...#
#..[]....#
#..@.....#
##########
				`,
			moves: "^",
			expected: `
##########
#........#
#..[]#...#
#.[][]...#
#..[]....#
#..@.....#
##########
`,
		},
		{
			grid: `
##########
#........#
#.[][][].#
#..[][]..#
#...[]...#
#...@....#
##########
`,
			moves: "^",
			expected: `
##########
#.[][][].#
#..[][]..#
#...[]...#
#...@....#
#........#
##########
`,
		},
	}

	for _, c := range cases {
		// grid, moves := parseInput([]byte(strings.TrimSpace(c.input)))
		grid := toGrid(strings.TrimSpace(c.grid))
		moveBoxes(grid, []byte(c.moves))

		g := bytes.Join(grid, []byte{'\n'})
		c.expected = strings.TrimSpace(c.expected)
		if c.expected != string(g) {
			t.Errorf("expected:\n%s\ngot:\n%s\n", c.expected, string(g))
		}
	}
}

func toGrid(input string) [][]byte {
	grid := [][]byte{}
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		grid = append(grid, []byte(l))
	}

	return grid
}
