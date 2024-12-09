package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 2858 {
		t.Errorf("expected 2858, got: %d", res)
	}
}
