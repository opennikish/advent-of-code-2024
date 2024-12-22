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

	valid := 0

	for _, design := range designs {
		if isValid(patterns, design) {
			valid++
		}
	}

	return valid
}

func isValid(patterns []string, design string) bool {
	if len(design) == 0 {
		return true
	}

	idx, found := sort.Find(len(patterns), func(j int) int {
		return int(design[0]) - int(patterns[j][0])
	})

	if !found {
		return false
	}

	for _, p := range patterns[idx:] {
		if p[0] != design[0] {
			break
		}
		if len(design) >= len(p) && design[:len(p)] == p {
			if isValid(patterns, design[len(p):]) {
				return true
			}
		}
	}

	return false
}

func parseInput(input []byte) ([]string, []string) {
	inputStr := strings.TrimSpace(string(input))
	parts := strings.SplitN(inputStr, "\n", 2)

	patterns := strings.Split(parts[0], ", ")
	designs := strings.Split(parts[1][1:], "\n")

	return patterns, designs
}
