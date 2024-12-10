package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	res := solve("in.txt")
	fmt.Println(res)
}

type Coord struct {
	i, j int
}

func solve(path string) int {
	bs, err := os.ReadFile(path)
	check(err)

	trailmap := bytes.Split(bytes.TrimSpace(bs), []byte{'\n'})
	trailheads := collectTrailHeads(trailmap)

	total := 0
	for _, coord := range trailheads {
		seen := map[Coord]bool{}
		total += countScore(trailmap, seen, coord.i, coord.j, '0'-1)
	}

	return total
}

func countScore(trailmap [][]byte, seen map[Coord]bool, row, col int, prev byte) int {
	n, m := len(trailmap), len(trailmap[0])

	if row < 0 || row >= n || col < 0 || col >= m {
		return 0
	}

	height := trailmap[row][col]

	if height == 'X' || height-1 != prev {
		return 0
	}

	if height == '9' {
		c := Coord{row, col}
		if seen[c] {
			return 0
		}
		seen[c] = true
		return 1
	}

	trailmap[row][col] = 'X'

	count := countScore(trailmap, seen, row, col+1, height) +
		countScore(trailmap, seen, row+1, col, height) +
		countScore(trailmap, seen, row, col-1, height) +
		countScore(trailmap, seen, row-1, col, height)

	trailmap[row][col] = height

	return count
}

func collectTrailHeads(trailmap [][]byte) []Coord {
	trailheads := []Coord{}

	for i := range trailmap {
		for j := range trailmap[0] {
			if trailmap[i][j] == '0' {
				trailheads = append(trailheads, Coord{i, j})
			}
		}
	}

	return trailheads
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
