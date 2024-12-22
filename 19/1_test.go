package main

import (
	"cmp"
	"slices"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	input := strings.TrimLeft(`
r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
`, "\n")

	res := solve([]byte(input))
	if res != 6 {
		t.Errorf("Got %d, expected 4", res)
	}
}

func TestIsValid(t *testing.T) {
	cases := []struct {
		design   string
		expected bool
	}{
		{
			"brwrr",
			true,
		},
		{
			"bggr",
			true,
		},
		{
			"gbbr",
			true,
		},
		{
			"rrbgbr",
			true,
		},
		{
			"ubwu",
			false,
		},
		{
			"bwurrg",
			true,
		},
		{
			"brgr",
			true,
		}, {
			"bbrgwb",
			false,
		},
	}

	patterns := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	slices.SortFunc(patterns, func(a, b string) int {
		return cmp.Compare(a[0], b[0])
	})

	for _, c := range cases {
		if isValid(patterns, c.design) != c.expected {
			t.Errorf("expected %t for %s", c.expected, c.design)
		}
	}
}
