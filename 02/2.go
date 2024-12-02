package main

import (
	"fmt"
)

func main() {
	fmt.Println(solve("in.txt"))
}

func solve(path string) int {
	levelsList := readLevels(path)

	safeCount := 0
	for _, levels := range levelsList {
		if safe(levels) {
			safeCount++
			continue
		}
		for i := 0; i < len(levels); i++ {
			if safe(append(append([]int(nil), levels[0:i]...), levels[i+1:]...)) {
				safeCount++
				break
			}
		}
	}

	return safeCount
}

func safe(levels []int) bool {
	dir := 1
	if levels[0] > levels[1] {
		dir = -1
	}
	for i := 1; i < len(levels); i++ {
		curr, prev := levels[i], levels[i-1]
		a, b := curr, prev
		if dir < 0 {
			a, b = b, a
		}
		if a <= b || a-b > 3 {
			return false
		}
	}

	return true
}
