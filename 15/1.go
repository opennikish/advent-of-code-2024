package main

import (
	"adventofcode2024/lib"
	"bytes"
	"fmt"
)

func main() {
	input, err := lib.GetInput(15)
	lib.Check(err)
	_ = input

	res := solve(input)
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

	gridBuilt := false

	for _, line := range lines {
		if len(line) == 0 {
			gridBuilt = true
			continue
		}
		if !gridBuilt {
			gl := append([]byte(nil), line...)
			grid = append(grid, gl)
		} else {
			moves = append(moves, line...)
		}
	}

	return grid, moves
}

func moveBoxes(grid [][]byte, moves []byte) {
	i, j := findRobotPos(grid)

	move := func(dir []int) {
		y := dir[0]
		x := dir[1]
		ni, nj := i, j
		canMove := false
		for grid[ni][nj] != '#' {
			ni += y
			nj += x
			if grid[ni][nj] == '.' {
				canMove = true
				break
			}
		}

		if !canMove {
			return
		}

		for ni != i || nj != j {
			grid[ni][nj], grid[ni-y][nj-x] = grid[ni-y][nj-x], grid[ni][nj]
			ni -= y
			nj -= x
		}

		i += y
		j += x
	}

	dirByMove := map[byte][]int{
		'^': []int{-1, 0},
		'>': []int{0, 1},
		'v': []int{1, 0},
		'<': []int{0, -1},
	}

	for _, m := range moves {
		dir := dirByMove[m]
		move(dir)
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
			if grid[i][j] == 'O' {
				total += i*100 + j
			}
		}
	}
	return total
}
