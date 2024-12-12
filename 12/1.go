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

type coord struct {
	i, j int
}

func solve(path string) int {
	bs, err := os.ReadFile(path)
	check(err)
	pmap := bytes.Split(bytes.TrimSpace(bs), []byte{'\n'})

	total := 0
	visited := map[coord]bool{}
	for i := range pmap {
		for j := range pmap[0] {
			if !visited[coord{i, j}] {
				total += calcPrice(pmap, i, j, visited)
			}
		}
	}

	return total
}

func calcPrice(pmap [][]byte, row, col int, visited map[coord]bool) int {
	plant := pmap[row][col]
	n, m := len(pmap), len(pmap[0])

	var dfs func(i, j int) (int, int)
	dfs = func(i, j int) (int, int) {
		if i < 0 || i >= n || j < 0 || j >= m {
			return 0, 0
		}

		c := coord{i, j}
		if visited[c] || pmap[i][j] != plant {
			return 0, 0
		}

		visited[c] = true

		area := 1
		peroid := calcPeroid(pmap, i, j)

		dirs := []coord{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, d := range dirs {
			a, p := dfs(i+d.i, j+d.j)
			area += a
			peroid += p
		}

		return area, peroid
	}

	a, p := dfs(row, col)
	return a * p
}

func calcPeroid(pmap [][]byte, i, j int) int {
	plant := pmap[i][j]
	n, m := len(pmap), len(pmap[0])

	peroid := 0
	dirs := []coord{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, d := range dirs {
		ni, nj := i+d.i, j+d.j
		if ni < 0 || ni >= n || nj < 0 || nj >= m || pmap[ni][nj] != plant {
			peroid++
			continue
		}
	}

	return peroid
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
