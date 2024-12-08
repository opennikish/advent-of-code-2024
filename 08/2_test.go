package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 34 {
		t.Errorf("expected 34, got: %d", res)
	}
}
