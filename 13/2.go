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

func solve(path string) int64 {
	bs, err := os.ReadFile(path)
	check(err)
	bs = bytes.TrimSpace(bs)
	machineInfos := strings.Split(string(bs), "\n\n")

	total := int64(0)
	for _, info := range machineInfos {
		total += findMinMovements(info)
	}
	return total
}

func findMinMovements(info string) int64 {
	lines := strings.Split(strings.TrimSpace(info), "\n")
	if len(lines) != 3 {
		panic("unexpected machine info line length")
	}

	parse := func(s string) (int64, int64) {
		rawNums := rNums.FindAllString(s, -1)
		if len(rawNums) != 2 {
			panic(fmt.Sprintf("unexpected raw nums length for A: %d", len(rawNums)))
		}
		return toInt64(rawNums[0]), toInt64(rawNums[1])
	}

	ax, ay := parse(lines[0])
	bx, by := parse(lines[1])
	targetX, targetY := parse(lines[2])
	targetX += 10000000000000
	targetY += 10000000000000

	a, b := solveEquation(ax, ay, bx, by, targetX, targetY)
	if a == -1 && b == -1 {
		return 0
	}

	return a*3 + b
}

func toInt64(s string) int64 {
	x, err := strconv.ParseInt(s, 10, 64)
	check(err)
	return x
}

// solveEquation uses Cramer's Rule to solve the equation
//
// Example:
// Button A: X+94, Y+34
// Button B: X+22, Y+67
// Prize: X=8400, Y=5400
//
// Example of equation:
// 94A + 22B = 8400
// 34A + 67B = 5400
//
// Could be written as:
// A*a_x + B*B_x = t_x
// A*a_y + B*b_y = t_y
//
// A = (t_x*b_y - t_y*b_x) / (a_x*b_y - a_y*b_x)
// B = (t_y*a_x - t_x*a_y) / (a_x*b_y - a_y*b_x)
func solveEquation(ax, ay, bx, by, targetX, targetY int64) (int64, int64) {
	det := ax*by - bx*ay
	a := (targetX*by - targetY*bx) / det
	b := (targetY*ax - targetX*ay) / det

	if ax*a+bx*b == targetX && ay*a+by*b == targetY {
		return a, b
	}

	return -1, -1
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
