package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	res := solve("in.txt")
	fmt.Println(res)
}

var rCommands = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
var rNums = regexp.MustCompile(`\d+`)

// xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
// ["mul(2,4)", mul(2,4), ..., "don't()", "mul(2,4)", "mul(2,4), ..., "do()", ..."]
func solve(path string) int {
	bs, err := os.ReadFile(path)
	check(err)
	instructions := string(bs)

	commands := rCommands.FindAllString(instructions, -1)
	sum := 0
	enabled := true

	for _, cmd := range commands {
		if cmd == "do()" {
			enabled = true
			continue
		}
		if cmd == "don't()" {
			enabled = false
			continue
		}

		if enabled {
			sum += doMul(cmd)
		}
	}

	return sum
}

func doMul(mul string) int {
	parts := rNums.FindAllString(mul, -1)

	a, err := strconv.Atoi(parts[0])
	check(err)
	b, err := strconv.Atoi(parts[1])
	check(err)

	return a * b
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
