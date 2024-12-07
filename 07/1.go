package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	res := solve("in.txt")
	fmt.Println(res)
}

func solve(path string) int {
	bs, err := os.ReadFile(path)
	check(err)
	s := strings.TrimSpace(string(bs))
	lines := strings.Split(s, "\n")

	sum_ := 0
	for _, l := range lines {
		parts := strings.Split(l, ": ")
		expected, err := strconv.Atoi(parts[0])
		check(err)

		values := toIntSlice(parts[1])

		if couldEquationBeTrue(expected, values) {
			sum_ += expected
		}
	}

	return sum_
}

func couldEquationBeTrue(expected int, values []int) bool {
	opCombs := combinations("+*", len(values)-1)

	for _, comb := range opCombs {
		res := values[0]
		for i := 1; i < len(values); i++ {
			op := comb[i-1]
			if op == '+' {
				res += values[i]
				continue
			}
			if op == '*' {
				res *= values[i]
				continue
			}
			panic("unexpected op: " + string(op))
		}

		if res == expected {
			return true
		}
	}

	return false
}

func combinations(alfpabet string, n int) []string {
	res := []string{}
	if n == 1 {
		for _, x := range alfpabet {
			res = append(res, string(x))
		}
		return res
	}

	subCombs := combinations(alfpabet, n-1)
	for _, x := range alfpabet {
		for _, sub := range subCombs {
			res = append(res, string(x)+sub)
		}
	}
	return res
}

func toIntSlice(rawValues string) []int {
	parts := strings.Split(rawValues, " ")
	values := make([]int, len(parts))
	for i, p := range parts {
		x, err := strconv.Atoi(p)
		check(err)
		values[i] = x
	}
	return values
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
