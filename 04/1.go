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
	matrix := readInput(path)
	n, m := len(matrix), len(matrix[0])
	count := 0
	dirs := [][]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for _, d := range dirs {
				if search(matrix, i, j, "XMAS", 0, d) {
					count++
				}
			}
		}
	}

	return count
}

func search(matrix [][]byte, row, col int, target string, ti int, dir []int) bool {
	n, m := len(matrix), len(matrix[0])
	if row < 0 || row >= n || col < 0 || col >= m {
		return false
	}

	if matrix[row][col] != target[ti] {
		return false
	}

	if ti == len(target)-1 {
		return true
	}

	return search(matrix, row+dir[0], col+dir[1], target, ti+1, dir)
}

func readInput(path string) [][]byte {
	bs, err := os.ReadFile(path)
	check(err)
	bs = bytes.TrimSpace(bs)
	matrix := bytes.Split(bs, []byte{'\n'})
	for i := 0; i < len(matrix); i++ {
		matrix[i] = bytes.TrimSpace(matrix[i])
	}
	return matrix
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
