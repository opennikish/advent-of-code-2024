package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 143 {
		t.Errorf("expected 143, got: %d", res)
	}
}
