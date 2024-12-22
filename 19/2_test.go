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
	if res != 16 {
		t.Errorf("Got %d, expected 16", res)
	}
}

func TestIsValid(t *testing.T) {
	cases := []struct {
		design   string
		expected int
	}{
		{
			"brwrr",
			2,
		},
		{
			"bggr",
			1,
		},
		{
			"gbbr",
			4,
		},
		{
			"rrbgbr",
			6,
		},
		{
			"ubwu",
			0,
		},
		{
			"bwurrg",
			1,
		},
		{
			"brgr",
			2,
		}, {
			"bbrgwb",
			0,
		},
	}

	patterns := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	slices.SortFunc(patterns, func(a, b string) int {
		return cmp.Compare(a[0], b[0])
	})

	cache := map[string]int{}

	for _, c := range cases {
		if countValid(patterns, cache, c.design) != c.expected {
			t.Errorf("expected %d for %s", c.expected, c.design)
		}
	}
}
