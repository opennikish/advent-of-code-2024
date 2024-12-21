package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var rNum = regexp.MustCompile(`\d+`)

type Regs struct {
	A, B, C int
}

type VM struct {
	Regs Regs
	out  []byte
}

var endlessLoop = errors.New("endless loop")

// Run runs program on the VM.
// Program consists of instruction code followed by operand.
//
// So, the program 0,1,2,3 would run the instruction whose opcode is 0 and pass it the operand 1,
// then run the instruction having opcode 2 and pass it the operand 3, then halt.
//
// VM registers (int): A, B, C
//
// There are two types of operands: literal (3-bit value itself) and combo (3-bit number [0-7]):
// 0-3 - literals values 0-3
// 4   - A's value
// 5   - B's value
// 6   - C's value
// 7   - reserved (should not appear in valid program)
//
// 3-bit instructions (opcode - opname; desc):
// 0 - adv; A = A / 2 ^ combo; A / 2^3; A / 2^B
// 1 - bxl (xor); B = B xor literal
// 2 - bst; B = combo % 8;
// 3 - jnz (jump not zero); if A == 0: do nothing; else: iPointer = literal
// 4 - bxc; B = B xor C (operand discarded)
// 5 - out; print(combo % 8); multiple values should be separated by commas
// 6 - bdv; B = A / 2 ^ combo;
// 7 - cdv; C = A / 2 ^ combo;
func (vm *VM) Run(program string) (string, error) {
	tokens := bytes.ReplaceAll([]byte(program), []byte{','}, nil)
	for i, t := range tokens {
		tokens[i] = t - '0'
	}

	opCombos := []bool{true, false, true, false, false, true, true, true}

	ip := 0
	for ip < len(tokens) {
		opCode := tokens[ip]
		ip++

		operand := int(tokens[ip])
		ip++
		if opCombos[opCode] {
			var err error
			operand, err = vm.resolveComboOperand(operand, ip)
			if err != nil {
				return "", err
			}
		}

		switch opCode {
		case 0:
			vm.Regs.A /= pow(2, operand)
		case 1:
			vm.Regs.B ^= operand
		case 2:
			vm.Regs.B = operand % 8
		case 3:
			if vm.Regs.A == 0 {
				continue
			}
			if operand%2 != 0 {
				return "", fmt.Errorf("jnz points to operand instead of op code at %d", ip)
			}
			ip = operand			
		case 4:
			vm.Regs.B ^= vm.Regs.C
		case 5:
			vm.print(byte(operand % 8))
		case 6:
			vm.Regs.B = vm.Regs.A / pow(2, operand)
		case 7:
			vm.Regs.C = vm.Regs.A / pow(2, operand)
		default:
			return "", fmt.Errorf("unexpected op code %d at %d", opCode, ip)
		}
	}

	return string(vm.out), nil
}

func (vm *VM) resolveComboOperand(operand int, ip int) (int, error) {
	if operand < 4 {
		return operand, nil
	}

	switch operand {
	case 4:
		return vm.Regs.A, nil
	case 5:
		return vm.Regs.B, nil
	case 6:
		return vm.Regs.C, nil
	}

	return 0, fmt.Errorf("unexpected operand %d at %d", operand, ip)
}

func (vm *VM) print(num byte) {
	if len(vm.out) > 0 {
		vm.out = append(vm.out, ',')
	}
	vm.out = append(vm.out, num+'0')
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

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
