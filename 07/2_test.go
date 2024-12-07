package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 11387 {
		t.Errorf("expected 11387, got: %d", res)
	}
}
