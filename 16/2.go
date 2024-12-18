package main

import (
	"adventofcode2024/lib"
	"bytes"
	"container/heap"
	"fmt"
	"math"
)

func main() {
	input, err := lib.GetInput(16)
	lib.Check(err)

	res := solve(input)
	fmt.Println(res)
}

type Face int

const (
	North Face = iota
	East
	South
	West
)

type dir struct {
	i, j int
}

type point struct {
	i, j int
	face Face
}

type Item struct {
	pos   point
	score int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].score < pq[j].score }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) { *pq = append(*pq, x.(Item)) }

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func solve(input []byte) int {
	bs := bytes.TrimSpace(input)
	maze := bytes.Split(bs, []byte{'\n'})

	n, m := len(maze), len(maze[0])

	row, col := n-2, 1
	if maze[row][col] != 'S' {
		panic("invalid start point: " + string(maze[row][col]))
	}

	start := point{row, col, East}

	dirs := []dir{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	dist := map[point]int{}
	for i := range maze {
		for j := range maze[0] {
			if maze[i][j] != '#' {
				for f := range dirs {
					dist[point{i, j, Face(f)}] = math.MaxInt
				}
			}
		}
	}
	dist[start] = 0

	paths := map[point][]point{}

	pq := PriorityQueue{}
	heap.Push(&pq, Item{pos: start, score: 0})
	heap.Init(&pq)

	visited := map[point]bool{}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(Item)

		if visited[it.pos] {
			continue
		}
		visited[it.pos] = true

		for i, dir := range dirs {
			f := Face(i)

			np := point{it.pos.i, it.pos.j, f}
			score := it.score + 1000

			if f == it.pos.face {
				np = point{it.pos.i + dir.i, it.pos.j + dir.j, it.pos.face}
				if np.i < 0 || np.i >= n || np.j < 0 && np.j >= m || maze[np.i][np.j] == '#' {
					continue
				}
				score = it.score + 1
			}

			if score <= dist[np] {
				if score < dist[np] {
					paths[np] = append([]point(nil), it.pos)
				} else { // equals
					paths[np] = append(paths[np], it.pos)
				}
				dist[np] = score
				heap.Push(&pq, Item{pos: np, score: score})
			}
		}
	}

	end := point{1, m - 2, East}
	if maze[end.i][end.j] != 'E' {
		panic("invalid end point")
	}

	return countBestPathsTiles(paths, end, maze)
}

func countBestPathsTiles(paths map[point][]point, end point, maze [][]byte) int {
	q := []point{end}
	all := map[dir]bool{}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		all[dir{curr.i, curr.j}] = true

		for _, p := range paths[curr] {
			q = append(q, p)
		}
	}

	for i := range maze {
		for j := range maze[0] {
			if all[dir{i, j}] {
				fmt.Print("O")
			} else {
				fmt.Print(string(maze[i][j]))
			}
		}
		fmt.Println()
	}

	return len(all)
}
