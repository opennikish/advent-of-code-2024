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

func solve(path string) int {
	instructions, err := os.ReadFile(path)
	check(err)

	rMuls := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	muls := rMuls.FindAllString(string(instructions), -1)

	rNums := regexp.MustCompile(`\d+`)

	sum := 0

	for _, mul := range muls {
		parts := rNums.FindAllString(mul, -1)

		a, err := strconv.Atoi(parts[0])
		check(err)
		b, err := strconv.Atoi(parts[1])
		check(err)

		sum += a * b
	}

	return sum
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
