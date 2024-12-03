package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test2.txt")
	if res != 48 {
		t.Errorf("expected 48, got: %d", res)
	}
}
