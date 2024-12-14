package main

import (
	"os"
	"testing"
)

func TestSolve(t *testing.T) {
	bs, err := os.ReadFile("test.txt")
	if err != nil {
		t.Errorf("read test file", err)
	}

	res := solve(bs)
	if res != 0 {
		t.Errorf("expected 0, got: %d", res)
	}
}
