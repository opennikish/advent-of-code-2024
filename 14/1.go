package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"adventofcode2024/lib"
)

func main() {
	input, err := lib.GetInput(14)
	lib.Check(err)

	res := solve(input, 101, 103, 100)
	fmt.Println(res)
}

type robot struct {
	x, y       int
	movX, movY int
}

type point struct {
	x, y int
}

var rNums = regexp.MustCompile(`[\-0-9]+`)

func solve(input []byte, width, height, sec int) int {
	robots := parseRobots(input)
	finalPosCounts := map[point]int{}

	transform := func(x, move, limit int) int {
		steps := abs((move * sec) % limit)
		if move > 0 {
			return (x + steps) % limit
		}

		x -= steps
		if x < 0 {
			x += limit
		}
		return x
	}

	for _, r := range robots {
		p := point{}
		p.x = transform(r.x, r.movX, width)
		p.y = transform(r.y, r.movY, height)

		finalPosCounts[p]++
	}

	skipX, skipY := width/2, height/2
	q1, q2, q3, q4 := 0, 0, 0, 0
	for pos, count := range finalPosCounts {
		if pos.y < skipY && pos.x < skipX {
			q1 += count
		}
		if pos.y < skipY && pos.x > skipX {
			q2 += count
		}
		if pos.y > skipY && pos.x < skipX {
			q3 += count
		}
		if pos.y > skipY && pos.x > skipX {
			q4 += count
		}
	}

	return q1 * q2 * q3 * q4
}

func parseRobots(input []byte) []robot {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	robots := make([]robot, len(lines))
	for i, l := range lines {
		nums := rNums.FindAllString(l, -1)
		if len(nums) != 4 {
			panic("corrupted robot: " + l)
		}
		robots[i] = robot{
			x:    toInt(nums[0]),
			y:    toInt(nums[1]),
			movX: toInt(nums[2]),
			movY: toInt(nums[3]),
		}
	}
	return robots
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	lib.Check(err)
	return n
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
