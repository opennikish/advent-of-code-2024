package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 14 {
		t.Errorf("expected 14, got: %d", res)
	}
}
