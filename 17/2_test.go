package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	input := strings.TrimLeft(`
Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0	
`, "\n")

	res, err := solve([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if res != 117440 {
		t.Errorf("expected: 117440, got: %d", res)
	}
}
