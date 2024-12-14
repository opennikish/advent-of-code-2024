package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 0 {
		t.Errorf("expected 0, got: %d", res)
	}
}
