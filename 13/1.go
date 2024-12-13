package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rNums = regexp.MustCompile(`\d+`)

func main() {
	res := solve("in.txt")
	fmt.Println(res)
}

func solve(path string) int {
	bs, err := os.ReadFile(path)
	check(err)
	bs = bytes.TrimSpace(bs)
	machineInfos := strings.Split(string(bs), "\n\n")

	total := 0
	for _, info := range machineInfos {
		total += findMinMovements(info)
	}
	return total
}

func findMinMovements(info string) int {
	lines := strings.Split(strings.TrimSpace(info), "\n")
	if len(lines) != 3 {
		panic("unexpected machine info line length")
	}

	parse := func(s string) (int, int) {
		rawNums := rNums.FindAllString(s, -1)
		if len(rawNums) != 2 {
			panic(fmt.Sprintf("unexpected raw nums length for A: %d", len(rawNums)))
		}
		return toInt(rawNums[0]), toInt(rawNums[1])
	}

	ax, ay := parse(lines[0])
	bx, by := parse(lines[1])
	targetX, targetY := parse(lines[2])

	a, b := solveEquation(ax, ay, bx, by, targetX, targetY)
	if a == -1 && b == -1 {
		return 0
	}

	return a*3 + b
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	check(err)
	return x
}

func solveEquation(ax, ay, bx, by, targetX, targetY int) (int, int) {
	for a := 1; a <= 100; a++ {
		for b := 1; b <= 100; b++ {
			x := ax*a + bx*b
			y := ay*a + by*b

			if x == targetX && y == targetY {
				return a, b
			}
		}
	}

	return -1, -1
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
