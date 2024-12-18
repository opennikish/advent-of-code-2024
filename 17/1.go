package main

import (
	"adventofcode2024/lib"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var rNum = regexp.MustCompile(`\d+`)

func main() {
	input, err := lib.GetInput(17)
	lib.Check(err)

	res, err := solve(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}

type Regs struct {
	A, B, C int
}

type VM struct {
	Regs     Regs
	out      []byte
	iPointer int
	tokens   []byte
}

// Run runs program on the VM.
// program consist of instruction code followed by combo operand,
// e.g. "0,1" means op code 0 and combo operand is 1.
//
// VM registers (int): A, B, C
//
// combo operands (3-bit number [0-7]):
// 0-3 - literals values 0-3
// 4 - A's value
// 5 - B's value
// 6 - C's value
// 7 - reserved (should not appear in valid program)
//
// 3-bit instructions (opcode - opname; desc):
// 0 - adv; A = A / 2 ^ combo; A / 2^3; A / 2^B
// 1 - bxl (xor); B = B xor combo
// 2 - bst; B = combo % 8;
// 3 - jnz (jump not zero); if A == 0: do nothing; else: iPointer = combo
// 4 - bxc; B = B xor C (also, read combo and discard it)
// 5 - out; print(combo % 8); multiple values should be separated by commas
// 6 - bdv; B = A / 2 ^ combo;
// 7 - cdv; C = A / 2 ^ combo;
func (vm *VM) Run(program string) string {
	tokens := bytes.ReplaceAll([]byte(program), []byte{','}, nil)
	_ = tokens

	return string(vm.out)
}

func (vm *VM) print(num byte) {
	if len(vm.out) > 0 {
		vm.out = append(vm.out, ',')
	}
	vm.out = append(vm.out, num+'0')
}

func solve(input []byte) (string, error) {
	regs, program, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("parse input: %w", err)
	}

	vm := &VM{Regs: regs}

	return vm.Run(program), nil
}

func parseInput(input []byte) (Regs, string, error) {
	var program string

	rowRegs := []int{}
	lines := bytes.Split(bytes.TrimSpace(input), []byte{'\n'})
	i := 0
	for ; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			break
		}

		s := rNum.FindString(string(lines[i]))
		if s == "" {
			return Regs{}, "", fmt.Errorf("invalid registor at line: %d", i+1)
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			return Regs{}, "", fmt.Errorf("strconv reg: %w", err)
		}
		rowRegs = append(rowRegs, v)
	}

	if len(rowRegs) != 3 {
		return Regs{}, "", fmt.Errorf("unexpected registor length: %v", rowRegs)
	}
	regs := Regs{
		A: rowRegs[0],
		B: rowRegs[1],
		C: rowRegs[2],
	}

	if i+1 >= len(lines) {
		return Regs{}, "", fmt.Errorf("corrupted input format")
	}

	program = string(lines[i+1])
	if len(program) == 0 {
		return Regs{}, "", fmt.Errorf("corrupted input format")
	}

	return regs, program[len("Program: "):], nil
}
