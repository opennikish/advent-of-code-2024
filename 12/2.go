package main

import (
	"bytes"
	"cmp"
	"fmt"
	"os"
	"slices"
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

type edge struct {
	dir, val int
}

func calcPrice(pmap [][]byte, row, col int, visited map[coord]bool) int {
	plant := pmap[row][col]
	n, m := len(pmap), len(pmap[0])

	colsByRow := map[int][]edge{} // for horizontal edges
	rowsByCol := map[int][]edge{} // for vertical edges

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= n || j < 0 || j >= m {
			return 0
		}

		c := coord{i, j}
		if visited[c] || pmap[i][j] != plant {
			return 0
		}

		visited[c] = true

		area := 1

		if i-1 < 0 || pmap[i-1][j] != plant {
			colsByRow[i] = append(colsByRow[i], edge{dir: 1, val: j})
		}
		if i+1 >= n || pmap[i+1][j] != plant {
			colsByRow[i] = append(colsByRow[i], edge{dir: -1, val: j})
		}
		if j-1 < 0 || pmap[i][j-1] != plant {
			rowsByCol[j] = append(rowsByCol[j], edge{dir: -1, val: i})
		}
		if j+1 >= m || pmap[i][j+1] != plant {
			rowsByCol[j] = append(rowsByCol[j], edge{dir: 1, val: i})
		}

		dirs := []coord{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, d := range dirs {
			area += dfs(i+d.i, j+d.j)
		}

		return area
	}

	area := dfs(row, col)
	sides := countSides(colsByRow) + countSides(rowsByCol)

	return area * sides
}

func countSides(locations map[int][]edge) int {
	count := 0
	for _, edges := range locations {
		slices.SortFunc(edges, func(a, b edge) int {
			return cmp.Or(
				cmp.Compare(a.dir, b.dir),
				cmp.Compare(a.val, b.val),
			)
		})
		prev := edge{dir: -2, val: -2}
		for _, e := range edges {
			if e.dir != prev.dir || e.val-1 != prev.val {
				count++
			}
			prev = e
		}
	}
	return count
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
