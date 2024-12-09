package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	res := solve("in.txt")
	fmt.Println(res)
}

func solve(path string) int {
	bs, err := os.ReadFile(path)
	check(err)
	fmap := bytes.TrimSpace(bs)
	fmt.Println(len(fmap))

	converted := convert(fmap)

	i, j := 0, len(converted)-1
	for i < j {
		if converted[i] != -1 {
			i++
			continue
		}
		if converted[j] == -1 {
			j--
			continue
		}

		converted[i], converted[j] = converted[j], converted[i]
		i++
		j--
	}

	return checksum(converted)
}

func checksum(fmap []int) int {
	s := 0
	for i, id := range fmap {
		if id == -1 {
			return s
		}
		s += id * i
	}
	return s
}

// 12345
// 0..111....22222
func convert(fmap []byte) []int {
	res := []int{}
	id := 0
	for i, b := range fmap {
		d := int(b - '0')

		var next int
		if (i+1)%2 == 0 {
			next = -1
		} else {
			next = id
			id++
		}

		res = append(res, slices.Repeat([]int{next}, d)...)
	}

	return res
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
