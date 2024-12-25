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

const StepLimit = 20

type point struct {
	i, j int
}

type pointpair struct {
	a, b point
}

func solve(input []byte, minCheatWin int) int {
	grid := parseInput(input)
	start := findCell(grid, S)
	path := traverse(grid, start)

	used := map[pointpair]bool{}
	cheats := 0
	for _, p := range path {
		cheats += bfs(grid, used, p, minCheatWin)
	}

	return cheats
}

func bfs(grid [][]int, used map[pointpair]bool, start point, minCheatWin int) int {
	dirs := []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	count := 0
	steps := 0
	seen := map[point]bool{start: true}
	q := []point{start}

	for len(q) > 0 && steps <= StepLimit {
		n := len(q)

		for _ = range n {
			p := q[0]
			q = q[1:]

			if p != start && grid[p.i][p.j] > 0 {
				if grid[p.i][p.j]-grid[start.i][start.j]-steps >= minCheatWin {
					from, to := pointpair{start, p}, pointpair{p, start}
					if !used[from] && !used[to] {
						used[from] = true
						used[to] = true

						count++
					}
				}
			}

			for _, dir := range dirs {
				cp := point{p.i + dir.i, p.j + dir.j}
				if cp.i < 0 || cp.j < 0 || cp.i >= len(grid) || cp.j >= len(grid[0]) || seen[cp] {
					continue
				}
				seen[cp] = true
				q = append(q, cp)
			}
		}

		steps++
	}

	return count
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
