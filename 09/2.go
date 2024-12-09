package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

type slot struct {
	i, size int
}

func main() {
	res := solve("in.txt")
	fmt.Println(res)
}

func print(fmap []int) {
	for _, x := range fmap {
		if x >= 0 {
			fmt.Printf("%d", x)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func solve(path string) int {
	bs, err := os.ReadFile(path)
	check(err)
	encodedFileMap := bytes.TrimSpace(bs)

	fmap := convert(encodedFileMap)
	// print(fmap)

	emptySlots := filterEmptySlots(fmap)

	for i := len(fmap) - 1; i >= 0; {
		if fmap[i] == -1 {
			i--
			continue
		}

		end, id := i, fmap[i]
		for i >= 0 && fmap[i] == id {
			i--
		}

		length := end - i
		for _, slot := range emptySlots {
			if slot.size < length || slot.i > i+1 {
				continue
			}

			copy(fmap[slot.i:slot.i+length], fmap[i+1:end+1])
			for j := i + 1; j <= end; j++ {
				fmap[j] = -1
			}

			slot.size -= length
			slot.i += length

			break
		}
	}

	// print(fmap)

	return checksum(fmap)
}

// [0,0,-1,-1,-1,-1,2,2,2,-1,-1]
func filterEmptySlots(fmap []int) []*slot {
	res := []*slot{}

	for i := 0; i < len(fmap); {
		x := fmap[i]
		if x != -1 {
			i++
			continue
		}

		start := i
		for i < len(fmap) && fmap[i] == -1 {
			i++
		}
		res = append(res, &slot{i: start, size: i - start})
	}

	return res
}

func checksum(fmap []int) int {
	s := 0
	for i, id := range fmap {
		if id == -1 {
			continue
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
