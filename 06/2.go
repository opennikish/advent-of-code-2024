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
	loopCache := map[string]bool{}

	for i >= 0 && i < len(labMap) && j >= 0 && j < len(labMap[0]) {
		if labMap[i][j] == '#' {
			// step back
			i -= dirs[d][0]
			j -= dirs[d][1]

			// turn right
			d = (d + 1) % len(dirs)
			continue
		}

		if hasLoop(labMap, i, j, dirs, d, loopCache) {
			// loopPos[i*len(labMap[0])+j] = true
			// loopCache[fmt.Sprintf("%d_%d", i, j)] = true
		}

		i += dirs[d][0]
		j += dirs[d][1]
	}

	return len(loopCache)
}

func hasLoop(labMap [][]byte, row, col int, dirs [][]int, startDir int, loopCache map[string]bool) bool {
	type rotate struct {
		i, j int
		d    int
	}

	seen := map[rotate]bool{}

	i, j := row, col

	ii, jj := i+dirs[startDir][0], j+dirs[startDir][1]
	if ii < 0 || ii >= len(labMap) || jj < 0 || jj >= len(labMap[0]) || labMap[ii][jj] == '#' || labMap[ii][jj] == '^' {
		return false
	}

	labMap[ii][jj] = '#'
	defer func() {
		labMap[ii][jj] = '.'
	}()

	d := startDir
	// moved := false
	// steps := 0
	for i >= 0 && i < len(labMap) && j >= 0 && j < len(labMap[0]) {
		// if steps > len(labMap)*len(labMap[0]) {
		// if steps > 1000000 {
		// 	// panic(fmt.Sprintf("endless %d,%d", ii, jj))
		// 	fmt.Printf("endless %d,%d\n", ii, jj)
		// 	return true
		// }
		// steps++

		// if moved && i == row && j == col && d == startDir {
		// 	return true
		// }
		// moved = true

		if labMap[i][j] == '#' {
			r := rotate{i: i, j: j, d: d}
			if seen[r] {
				loopCache[fmt.Sprintf("%d_%d", ii, jj)] = true
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
