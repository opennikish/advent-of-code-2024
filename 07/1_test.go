package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 3749 {
		t.Errorf("expected 3749, got: %d", res)
	}
}
