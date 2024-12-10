package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 81 {
		t.Errorf("expected 81, got: %d", res)
	}
}
