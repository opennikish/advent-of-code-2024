package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 4 {
		t.Errorf("expected 4, got %d", res)
	}
}
