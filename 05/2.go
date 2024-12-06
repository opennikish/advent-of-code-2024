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
	prevsByNext, nextsByPrevs, updates := parseInput(path)
	middleSum := 0
	for _, upd := range updates {
		if !isValidUpdate(upd, prevsByNext) {
			fixedUpd := fixOrder(upd, nextsByPrevs)
			m := len(fixedUpd) / 2
			middleSum += fixedUpd[m]
		}
	}
	return middleSum
}

func fixOrder(upd []int, nextsByPrev map[int]map[int]bool) []int {
	g := map[int][]int{}

	for _, page := range upd {
		g[page] = findIntersection(upd, nextsByPrev[page])
	}

	visited := map[int]bool{}
	sorted := make([]int, len(g))
	counter := 0
	for page := range g {
		if !visited[page] {
			dfs(g, visited, sorted, &counter, page, page)
		}
	}

	return sorted
}

func dfs(g map[int][]int, visited map[int]bool, sorted []int, counter *int, v int, from int) {
	visited[v] = true

	for _, u := range g[v] {
		if !visited[u] {
			dfs(g, visited, sorted, counter, u, v)
		}
	}

	sorted[len(sorted)-1-*counter] = v
	*counter++
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

func findIntersection(upd []int, set map[int]bool) []int {
	intersec := []int{}
	for _, page := range upd {
		if set[page] {
			intersec = append(intersec, page)
		}
	}
	return intersec
}

func parseInput(path string) (map[int]map[int]bool, map[int]map[int]bool, [][]int) {
	bs, err := os.ReadFile(path)
	check(err)
	input := string(bs)
	lines := strings.Split(input, "\n")

	prevsByNext := map[int]map[int]bool{}
	nextsByPrev := map[int]map[int]bool{}
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

			nexts, ok := nextsByPrev[prev]
			if !ok {
				nexts = map[int]bool{}
				nextsByPrev[prev] = nexts
			}
			nexts[next] = true

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

	return prevsByNext, nextsByPrev, updates
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
