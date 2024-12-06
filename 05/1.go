package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	res := solve("in.txt")
	fmt.Println(res)
}

func solve(path string) int {
	prevsByNext, updates := parseInput(path)
	middleSum := 0
	for _, upd := range updates {
		if isValidUpdate(upd, prevsByNext) {
			m := len(upd) / 2
			middleSum += upd[m]
		}
	}
	return middleSum
}

func isValidUpdate(upd []int, prevsByNext map[int]map[int]bool) bool {
	currs := []int{}
	isValid := true
	for _, next := range upd {
		prevs := findIntersection(upd, prevsByNext[next])

		if !containsAll(currs, prevs) {
			isValid = false
			break
		}

		currs = append(currs, next)
	}
	return isValid
}

func containsAll(currs, prevs []int) bool {
	for _, p := range prevs {
		if !slices.Contains(currs, p) {
			return false
		}
	}
	return true
}

func findIntersection(upd []int, allPrevs map[int]bool) []int {
	prevs := []int{}
	for _, page := range upd {
		if allPrevs[page] {
			prevs = append(prevs, page)
		}
	}
	return prevs
}

func parseInput(path string) (map[int]map[int]bool, [][]int) {
	bs, err := os.ReadFile(path)
	check(err)
	input := string(bs)
	lines := strings.Split(input, "\n")

	prevsByNext := map[int]map[int]bool{}
	updates := [][]int{}
	seenSep := false
	for _, l := range lines {
		if !seenSep {
			if l == "" {
				seenSep = true
				continue
			}

			parts := strings.Split(l, "|")
			if len(parts) != 2 {
				panic("invalid rule")
			}
			prev, err := strconv.Atoi(parts[0])
			check(err)
			next, err := strconv.Atoi(parts[1])

			prevs, ok := prevsByNext[next]
			if !ok {
				prevs = map[int]bool{}
				prevsByNext[next] = prevs
			}

			prevs[prev] = true

			continue
		}

		if l == "" {
			break
		}

		update := []int{}
		parts := strings.Split(l, ",")
		for _, p := range parts {
			x, err := strconv.Atoi(p)
			check(err)
			update = append(update, x)
		}

		updates = append(updates, update)
	}

	return prevsByNext, updates
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
