package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	input := []byte(strings.TrimLeft(`
	past input	
	`, "\n"))
	expected := 0
	res := solve(input)
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
