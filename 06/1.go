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

	i, j := findGuardPos(labMap)
	d, dirs := 0, [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	steps := 0
	for i >= 0 && i < len(labMap) && j >= 0 && j < len(labMap[0]) {
		if labMap[i][j] == '#' {
			// step back
			i -= dirs[d][0]
			j -= dirs[d][1]

			// turn right
			d = (d + 1) % len(dirs)
			continue
		}

		if labMap[i][j] != 'X' {
			steps++
		}
		labMap[i][j] = 'X'

		i += dirs[d][0]
		j += dirs[d][1]
	}

	return steps
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
