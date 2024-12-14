package main

import (
	"fmt"

	"adventofcode2024/lib"
)

func main() {
	input, err := lib.GetInput(14)
	lib.Check(err)

	res := solve(input)
	fmt.Println(res)
}

func solve(input []byte) int {
	return 0
}
