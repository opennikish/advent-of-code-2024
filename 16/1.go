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

type point struct {
	i, j int
}

type Item struct {
	pos   point
	score int
	face  Face
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

	start := point{row, col}

	dist := map[point]int{}
	for i := range maze {
		for j := range maze[0] {
			if maze[i][j] != '#' {
				dist[point{i, j}] = math.MaxInt
			}
		}
	}
	dist[start] = 0

	pq := PriorityQueue{}
	heap.Push(&pq, Item{pos: start, score: 0, face: East})
	heap.Init(&pq)

	dirs := []point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(Item)

		// vertex could be added multiple times.
		// process only one with lowest distance (score)
		if it.score > dist[it.pos] {
			continue
		}

		for i, d := range dirs {
			np := point{it.pos.i + d.i, it.pos.j + d.j}
			if np.i < 0 || np.i >= n || np.j < 0 || np.j >= m || maze[np.i][np.j] == '#' {
				continue
			}

			f := Face(i)
			score := it.score + 1

			if f != it.face {
				score += 1000
			}

			if score < dist[np] {
				dist[np] = score
				heap.Push(&pq, Item{pos: np, score: score, face: f})
			}
		}
	}

	end := point{1, m - 2}
	if maze[end.i][end.j] != 'E' {
		panic("invalid end point")
	}

	return dist[end]
}
