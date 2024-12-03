package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 161 {
		t.Errorf("expected 161, got: %d", res)
	}
}
