package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 41 {
		t.Errorf("expected 41, got: %d", res)
	}
}
