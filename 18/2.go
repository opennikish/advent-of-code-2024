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

func solve(input []byte, pointSize, gridSize int) string {
	points := parseInput(input)

	grid := make([][]byte, gridSize+1)
	for i, _ := range grid {
		grid[i] = make([]byte, gridSize+1)
	}

	for i := range pointSize {
		p := points[i]
		grid[p.y][p.x] = '#'
	}

	for i := pointSize; i < len(points); i++ {
		p := points[i]
		grid[p.y][p.x] = '#'
		if bfs(grid) == -1 {
			return fmt.Sprintf("%d,%d", p.x, p.y)
		}
	}

	return "not found"
}

func bfs(grid [][]byte) int {
	type item struct {
		steps int
		p     point
	}

	start := point{x: 0, y: 0}
	q := []item{{steps: 0, p: start}}
	visited := map[point]bool{start: true}
	dirs := []point{{y: 0, x: 1}, {y: 1, x: 0}, {y: 0, x: -1}, {y: -1, x: 0}}
	n := len(grid)

	for len(q) > 0 {
		it := q[0]
		q = q[1:]

		if it.p.y == n-1 && it.p.x == n-1 {
			return it.steps
		}

		for _, d := range dirs {
			ni, nj := it.p.y+d.y, it.p.x+d.x
			if ni >= 0 && ni < n && nj >= 0 && nj < n && !visited[point{x: nj, y: ni}] && grid[ni][nj] != '#' {
				q = append(q, item{steps: it.steps + 1, p: point{x: nj, y: ni}})
				visited[point{x: nj, y: ni}] = true
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
