package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func readLevels(path string) [][]int {
	input, err := os.ReadFile(path)
	check(err)
	lines := strings.Split(string(input), "\n")

	levelsList := [][]int{}
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		parts := strings.Split(l, " ")
		nums := []int{}
		for _, x := range parts {
			num, err := strconv.Atoi(x)
			check(err)
			nums = append(nums, num)
		}

		levelsList = append(levelsList, nums)
	}

	return levelsList
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
