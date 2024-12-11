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
		total += countScore(trailmap, coord)
	}

	return total
}

func countScore(trailmap [][]byte, pos Coord) int {
	type Item struct {
		prev byte
		pos  Coord
	}
	n, m := len(trailmap), len(trailmap[0])

	q := []Item{{prev: '0' - 1, pos: pos}}
	seen := map[Coord]bool{}
	endCount := 0

	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		i, j := item.pos.i, item.pos.j

		if i < 0 || i >= n || j < 0 || j >= m {
			continue
		}

		height := trailmap[i][j]
		if seen[item.pos] || height-1 != item.prev {
			continue
		}
		seen[item.pos] = true

		if height == '9' {
			endCount++
		}

		q = append(q, Item{prev: height, pos: Coord{i, j + 1}},
			Item{prev: height, pos: Coord{i + 1, j}},
			Item{prev: height, pos: Coord{i, j - 1}},
			Item{prev: height, pos: Coord{i - 1, j}},
		)
	}

	return endCount
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
