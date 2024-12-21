package main

import (
	"adventofcode2024/lib"
	"fmt"
	"os"
)

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

func solve(input []byte) (string, error) {
	regs, program, err := parseInput(input)
	if err != nil {
		return "", fmt.Errorf("parse input: %w", err)
	}

	vm := &VM{Regs: regs}

	return vm.Run(program)
}
