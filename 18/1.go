package main

import (
	"adventofcode2024/lib"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input, err := lib.GetInput(18)
	lib.Check(err)

	res := solve(input, 1024, 70)
	fmt.Println(res)
}

type point struct {
	x, y int
}

func solve(input []byte, pointSize, gridSize int) int {
	points := parseInput(input)
	points = points[0:pointSize]

	grid := make([][]byte, gridSize+1)
	for i, _ := range grid {
		grid[i] = make([]byte, gridSize+1)
	}

	for _, p := range points {
		grid[p.y][p.x] = '#'
	}

	return bfs(grid)
}

func bfs(grid [][]byte) int {
	type item struct {
		steps int
		p     point
	}
	n := len(grid)
	q := []item{{steps: 0, p: point{x: 0, y: 0}}}
	grid[0][0] = 'O'
	dirs := []point{{y: 0, x: 1}, {y: 1, x: 0}, {y: 0, x: -1}, {y: -1, x: 0}}

	for len(q) > 0 {
		it := q[0]
		q = q[1:]

		if it.p.y == n-1 && it.p.x == n-1 {
			return it.steps
		}

		for _, d := range dirs {
			ni, nj := it.p.y+d.y, it.p.x+d.x
			if ni >= 0 && ni < n && nj >= 0 && nj < n && grid[ni][nj] != 'O' && grid[ni][nj] != '#' {
				q = append(q, item{steps: it.steps + 1, p: point{x: nj, y: ni}})
				grid[ni][nj] = 'O'
			}
		}
	}

	return -1
}

func parseInput(input []byte) []point {
	res := []point{}
	lines := bytes.Split(bytes.TrimSpace(input), []byte{'\n'})
	for _, l := range lines {
		parts := strings.Split(string(l), ",")
		if len(parts) != 2 {
			panic("invalid point line")
		}
		x, err := strconv.Atoi(parts[0])
		check(err)
		y, err := strconv.Atoi(parts[1])
		res = append(res, point{x: x, y: y})
	}
	return res
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
