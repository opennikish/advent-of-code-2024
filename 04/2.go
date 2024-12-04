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

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if contains(matrix, i, j) {
				count++
			}
		}
	}

	return count
}

func contains(matrix [][]byte, row, col int) bool {
	n, m := len(matrix), len(matrix[0])

	check := func(i, j int, target string, dir int) bool {
		if i+len(target)-1 >= n {
			return false
		}

		maxCol := j + dir*(len(target)-1)
		if maxCol < 0 || maxCol >= m {
			return false
		}

		for k := 0; k < len(target); k++ {
			if matrix[i][j] != target[k] {
				return false
			}
			i++
			j += dir
		}

		return true
	}

	return (check(row, col, "MAS", 1) || check(row, col, "SAM", 1)) &&
		(check(row, col+2, "MAS", -1) || check(row, col+2, "SAM", -1))
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
