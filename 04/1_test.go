package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 18 {
		t.Errorf("expected 18, got: %d", res)
	}
}
