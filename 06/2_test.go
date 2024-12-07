package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 6 {
		t.Errorf("expected 6, got: %d", res)
	}
}
