package main

import (
	"adventofcode2024/lib"
	"fmt"
)

func main() {
	input, err := lib.GetInput(0)
	lib.Check(err)

	res := solve(input)
	fmt.Println(res)
}

func solve(input []byte) int {
	return 0
}
