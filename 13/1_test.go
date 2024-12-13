package main

import "testing"

func TestSolve(t *testing.T) {
	res := solve("test.txt")
	if res != 480 {
		t.Errorf("expected 480, got: %d", res)
	}
}

func TestSolveEquation(t *testing.T) {
	a, b := solveEquation(94, 34, 22, 67, 8400, 5400)
	if a != 80 || b != 40 {
		t.Errorf("expected a=80 and b=40, got a=%d and b=%d", a, b)
	}
}

func TestFindMinMovements(t *testing.T) {
	res := findMinMovements(`
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400
	`)
	if res != 280 {
		t.Errorf("expected 280, got %d", res)
	}
}
