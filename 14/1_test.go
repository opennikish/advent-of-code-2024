package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	robots := `
p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
`
	res := solve([]byte(robots), 11, 7, 100)
	if res != 12 {
		t.Errorf("expected 12, got: %d", res)
	}
}
