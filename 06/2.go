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

func solve(path string) int {
	bs, err := os.ReadFile(path)
	check(err)
	bs = bytes.TrimSpace(bs)
	labMap := bytes.Split(bs, []byte{'\n'})

	row, col := findGuardPos(labMap)
	d, dirs := 0, [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	fillRoute(labMap, row, col, dirs, d)

	count := 0
	for i := 0; i < len(labMap); i++ {
		for j := 0; j < len(labMap[0]); j++ {
			if labMap[i][j] == 'X' {
				labMap[i][j] = '#'
				if hasLoop(labMap, row, col, dirs, d) {
					count++
				}
				labMap[i][j] = 'X'
			}
		}
	}

	return count
}

func fillRoute(labMap [][]byte, row, col int, dirs [][]int, d int) {
	i, j := row, col
	for i >= 0 && i < len(labMap) && j >= 0 && j < len(labMap[0]) {
		if labMap[i][j] == '#' {
			// step back
			i -= dirs[d][0]
			j -= dirs[d][1]

			// turn right
			d = (d + 1) % len(dirs)
			continue
		}

		if labMap[i][j] != '^' {
			labMap[i][j] = 'X'
		}

		i += dirs[d][0]
		j += dirs[d][1]
	}
}

func hasLoop(labMap [][]byte, row, col int, dirs [][]int, startDir int) bool {
	type rotation struct {
		i, j int
		d    int
	}

	seen := map[rotation]bool{}
	i, j := row, col
	d := startDir

	for i >= 0 && i < len(labMap) && j >= 0 && j < len(labMap[0]) {
		if labMap[i][j] == '#' {
			r := rotation{i: i, j: j, d: d}
			if seen[r] {
				return true
			}
			seen[r] = true

			// step back
			i -= dirs[d][0]
			j -= dirs[d][1]

			// turn right
			d = (d + 1) % len(dirs)
			continue
		}

		i += dirs[d][0]
		j += dirs[d][1]
	}

	return false
}

func findGuardPos(labMap [][]byte) (int, int) {
	for i := 0; i < len(labMap); i++ {
		for j := 0; j < len(labMap[0]); j++ {
			if labMap[i][j] == '^' {
				return i, j
			}
		}
	}
	panic("guard not found")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
