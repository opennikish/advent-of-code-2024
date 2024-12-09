package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 1928 {
		t.Errorf("expected 1928, got: %d", res)
	}
}
