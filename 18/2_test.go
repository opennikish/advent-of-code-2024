package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	input := strings.TrimLeft(`
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`, "\n")

	res := solve([]byte(input), 12, 6)
	if res != "6,1" {
		t.Errorf("expected 6,1, got: %s", res)
	}
}
