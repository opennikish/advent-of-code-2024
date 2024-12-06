package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 123 {
		t.Errorf("expected 123, got: %d", res)
	}
}
