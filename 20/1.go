package main

import (
	"adventofcode2024/lib"
	"bytes"
	"fmt"
)

func main() {
	input, err := lib.GetInput(20)
	lib.Check(err)

	res := solve(input, 100)
	fmt.Println(res)
}

const S = -2
const E = -3

type point struct {
	i, j int
}

func solve(input []byte, minCheatWin int) int {
	grid := parseInput(input)
	start, _ := findCell(grid, S), findCell(grid, E)
	path := traverse(grid, start)

	cheats := 0
	dirs := []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, p := range path {
		for _, d := range dirs {
			cp := point{p.i + d.i, p.j + d.j}
			if cp.i < 0 || cp.j < 0 || cp.i >= len(grid) || cp.j >= len(grid[0]) || grid[cp.i][cp.j] != -1 { // should go though walls
				continue
			}

			cp.i += d.i
			cp.j += d.j
			if cp.i < 0 || cp.j < 0 || cp.i >= len(grid) || cp.j >= len(grid[0]) || (grid[cp.i][cp.j] <= 0) { // should be part not seen cell on path
				continue
			}

			if abs(grid[cp.i][cp.j]-grid[p.i][p.j])-2 >= minCheatWin {
				fmt.Println(grid[p.i][p.j], "---", grid[cp.i][cp.j], "=>", abs(grid[cp.i][cp.j]-grid[p.i][p.j])-2)
				cheats++
			}
		}
		grid[p.i][p.j] = -4 // each cheat is unique
	}

	return cheats
}

func traverse(grid [][]int, start point) []point {
	dirs := []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	path := []point{}

	seen := map[point]bool{}

	var move func(i, j, count int)
	move = func(i, j, count int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == -1 || seen[point{i, j}] {
			return
		}
		seen[point{i, j}] = true
		path = append(path, point{i, j})
		grid[i][j] = count

		for _, dir := range dirs {
			move(i+dir.i, j+dir.j, count+1)
		}
	}

	move(start.i, start.j, 0)

	return path
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findCell(grid [][]int, target int) point {
	for i, row := range grid {
		for j, cell := range row {
			if cell == target {
				return point{i, j}
			}
		}
	}
	panic(fmt.Sprintf("cell %d not found", target))
}

func parseInput(input []byte) [][]int {
	grid := [][]int{}

	lines := bytes.Split(bytes.TrimSpace(input), []byte("\n"))
	for _, line := range lines {
		row := []int{}
		for _, b := range line {
			switch b {
			case '.':
				row = append(row, 0)
			case '#':
				row = append(row, -1)
			case 'S':
				row = append(row, S)
			case 'E':
				row = append(row, E)
			}
		}
		grid = append(grid, row)
	}
	return grid
}
