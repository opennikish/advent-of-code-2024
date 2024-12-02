package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 2 {
		t.Fail()
	}
}
