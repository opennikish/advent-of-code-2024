package main

import (
	"strings"
	"testing"
)

func TestRunVM(t *testing.T) {
	cases := []struct {
		regs         Regs
		expectedRegs *Regs // if nil don't check
		program      string
		output       string // if empty don't check
	}{
		{
			regs:         Regs{C: 9},
			expectedRegs: &Regs{B: 1},
			program:      "2,6",
			output:       "",
		},
		{
			regs:         Regs{A: 10},
			expectedRegs: nil,
			program:      "5,0,5,1,5,4",
			output:       "0,1,2",
		},
		{
			regs:         Regs{A: 2024},
			expectedRegs: &Regs{A: 0},
			program:      "0,1,5,4,3,0",
			output:       "4,2,5,6,7,7,7,7,3,1,0 ",
		},
		{
			regs:         Regs{B: 29},
			expectedRegs: &Regs{B: 26},
			program:      "1,7",
			output:       "",
		},
		{
			regs:         Regs{B: 2024, C: 43690},
			expectedRegs: &Regs{B: 44354},
			program:      "4,0",
			output:       "",
		},
		{
			regs:         Regs{A: 729},
			expectedRegs: nil,
			program:      "0,1,5,4,3,0",
			output:       "4,6,3,5,6,3,5,2,1,0",
		},
	}

	for _, c := range cases {
		vm := &VM{Regs: c.regs}

		out := vm.Run(c.program)
		if out != c.output {
			t.Errorf("expected: %s, got: %s", c.output, out)
		}
		if c.expectedRegs != nil {
			if vm.Regs != *c.expectedRegs {
				t.Errorf("expected: %+v, got: %+v", *c.expectedRegs, vm.Regs)
			}
		}
	}
}

func TestParseInput(t *testing.T) {
	input := strings.TrimLeft(`
Register A: 729
Register B: 20
Register C: 24

Program: 0,1,5,4,3,0
`, "\n")

	regs, program, err := parseInput([]byte(input))
	if err != nil {
		t.Fatal(err)
	}

	expectedRegs := Regs{A: 729, B: 20, C: 24}
	if regs != expectedRegs {
		t.Fatalf("expected: %+v, got: %+v", expectedRegs, regs)
	}

	expectedProgram := "0,1,5,4,3,0"
	if program != expectedProgram {
		t.Fatalf("expected: %s, got: %s", expectedProgram, program)
	}
}
