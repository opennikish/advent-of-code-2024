package main

import (
	"adventofcode2024/lib"
	"cmp"
	"fmt"
	"slices"
	"sort"
	"strings"
)

func main() {
	input, err := lib.GetInput(19)
	lib.Check(err)

	res := solve(input)
	fmt.Println(res)
}

func solve(input []byte) int {
	patterns, designs := parseInput(input)

	slices.SortFunc(patterns, func(a, b string) int {
		return cmp.Compare(a[0], b[0])
	})

	allValid := 0
	cache := map[string]int{}

	for _, design := range designs {
		allValid += countValid(patterns, cache, design)
	}

	return allValid
}

func countValid(patterns []string, cache map[string]int, design string) int {
	if len(design) == 0 {
		return 1
	}

	if count, ok := cache[design]; ok {
		return count
	}

	idx, found := sort.Find(len(patterns), func(j int) int {
		return int(design[0]) - int(patterns[j][0])
	})

	if !found {
		return 0
	}

	count := 0
	for _, p := range patterns[idx:] {
		if p[0] != design[0] {
			break
		}
		if len(design) >= len(p) && design[:len(p)] == p {
			count += countValid(patterns, cache, design[len(p):])
		}
	}

	cache[design] = count

	return count
}

func parseInput(input []byte) ([]string, []string) {
	inputStr := strings.TrimSpace(string(input))
	parts := strings.SplitN(inputStr, "\n", 2)

	patterns := strings.Split(parts[0], ", ")
	designs := strings.Split(parts[1][1:], "\n")

	return patterns, designs
}
