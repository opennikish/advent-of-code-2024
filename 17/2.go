package main

import (
	"adventofcode2024/lib"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

var notFound = errors.New("not found")

func main() {
	input, err := lib.GetInput(17)
	lib.Check(err)

	res, err := solve(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}

func solve(input []byte) (int, error) {
	_, program, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	res := bfs(program)

	if len(res) == 0 {
		return 0, notFound
	}

	return slices.Min(res), nil
}

func bfs(program string) []int {
	tokens := toTokens(program)
	n := len(tokens)

	type item struct {
		i, res int
	}

	q := []item{{i: n - 1, res: 0}}
	res := []int{}

	for len(q) > 0 {
		it := q[0]
		q = q[1:]

		if it.i < 0 {
			res = append(res, it.res)
			continue
		}

		for x := range 8 { // b/c 3 each output number is 3-bit
			a := (it.res << 3) | x

			vm := &VM{}
			vm.Regs.A = a
			out, _ := vm.Run(program)

			if out == toProgram(tokens[it.i:]) {
				fmt.Printf("ok, a: %d, bin: %b\n", a, a)
				q = append(q, item{i: it.i - 1, res: a})
			}
		}
	}

	return res
}

func toTokens(program string) []int {
	nums := strings.Split(strings.TrimSpace(program), ",")
	tokens := make([]int, len(nums))
	for i, s := range nums {
		tokens[i] = int(s[0] - '0')
	}
	return tokens
}

func toProgram(tokens []int) string {
	strs := make([]string, len(tokens))
	for i, t := range tokens {
		strs[i] = string(byte(t) + '0')
	}
	return strings.Join(strs, ",")
}
