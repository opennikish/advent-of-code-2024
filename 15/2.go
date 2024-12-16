package main

import (
	"adventofcode2024/lib"
	"bytes"
	"fmt"
	"os"
	"slices"
	"time"
)

type point struct {
	i, j int
}
type boxPos struct {
	left  point
	right point
}

func main() {
	var input []byte
	var err error
	if len(os.Args) > 1 {
		input, err = os.ReadFile(os.Args[1])
	} else {
		input, err = lib.GetInput(15)
	}
	lib.Check(err)

	res := solve([]byte(input))
	fmt.Println(res)
}

func solve(input []byte) int {
	grid, moves := parseInput(input)
	moveBoxes(grid, moves)
	return calcPointsSum(grid)
}

func parseInput(input []byte) ([][]byte, []byte) {
	lines := bytes.Split(input, []byte{'\n'})

	grid, moves := [][]byte{}, []byte{}

	i := 0
	for ; i < len(lines) && len(lines[i]) > 0; i++ {
		next := []byte{}
		for j := range lines[0] {
			switch lines[i][j] {
			case '.':
				next = append(next, '.', '.')
			case '#':
				next = append(next, '#', '#')
			case 'O':
				next = append(next, '[', ']')
			case '@':
				next = append(next, '@', '.')
			}
		}
		grid = append(grid, next)
	}
	i++ // skip empty line

	for ; i < len(lines); i++ {
		moves = append(moves, lines[i]...)
	}

	return grid, moves
}

func moveBoxes(grid [][]byte, moves []byte) {
	row, col := findRobotPos(grid)

	moveHorizontally := func(dir int) {
		nj := col
		canMove := false
		for grid[row][nj] != '#' {
			nj += dir
			if grid[row][nj] == '.' {
				canMove = true
				break
			}
		}

		if !canMove {
			return
		}

		for nj != col {
			grid[row][nj], grid[row][nj-dir] = grid[row][nj-dir], grid[row][nj]
			nj -= dir
		}

		col += dir
	}

	moveVertically := func(dir int) {
		rows := [][]point{}
		curr := []point{{row, col}}

		for len(curr) > 0 { // if len == 0 we reached only-dots line
			rows = append(rows, curr)
			next := []point{}

			for _, p := range curr {
				if grid[p.i+dir][p.j] == '#' {
					return
				}

				if grid[p.i+dir][p.j] == '.' {
					continue
				}

				// []
				// []
				// so, add only half of the upper box
				if grid[p.i][p.j] == grid[p.i+dir][p.j] {
					next = append(next, point{p.i + dir, p.j})
					continue
				}

				// [][]
				//  []
				// so, add whole upper box
				left := p.j
				if grid[p.i+dir][left] == ']' {
					left = p.j - 1
				}
				upperLeft := point{p.i + dir, left}

				// upper left box could be added by previous point (]) in a row:
				// #..[]....# <-  next row
				// #.[][]...# <-- curr row, right box is `p`
				if len(next) < 2 || next[len(next)-2] != upperLeft {
					next = append(next, upperLeft)
					next = append(next, point{p.i + dir, left + 1})
				}
			}

			curr = next
		}

		slices.Reverse(rows)
		for _, row := range rows {
			for _, p := range row {
				grid[p.i+dir][p.j], grid[p.i][p.j] = grid[p.i][p.j], grid[p.i+dir][p.j]
			}
		}

		row += dir
	}

	render(grid, moves, len(moves))
	time.Sleep(1 * time.Second)

	for i, m := range moves {
		switch m {
		case '^':
			moveVertically(-1)
		case '>':
			moveHorizontally(1)
		case 'v':
			moveVertically(1)
		case '<':
			moveHorizontally(-1)
		}

		render(grid, moves, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func render(grid [][]byte, moves []byte, curr int) {
	fmt.Print("\033[H\033[2J") // clear screen
	fmt.Println("ROBO WAREHOUSE\n")
	m := append([]byte(nil), moves...)
	if curr < len(m) {
		m = append(m[:curr], append(append([]byte(nil), '|', moves[curr], '|'), m[curr+1:]...)...)
	}

	i, width, aligned := 0, len(grid[0]), []byte{}
	for i < len(m) {
		limit := i + width
		if limit > len(m) {
			limit = len(m)
		}
		aligned = append(aligned, m[i:limit]...)
		aligned = append(aligned, '\n')
		i += width
	}

	fmt.Println(string(aligned))
	for i := range grid {
		fmt.Println(string(grid[i]))
	}
}

func findRobotPos(grid [][]byte) (int, int) {
	n, m := len(grid), len(grid[0])
	for i := range n {
		for j := range m {
			if grid[i][j] == '@' {
				return i, j
			}
		}
	}
	panic("cannot find robot position")
}

func calcPointsSum(grid [][]byte) int {
	total := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == '[' {
				total += i*100 + j
			}
		}
	}
	return total
}
