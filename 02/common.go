package main

import (
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
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
