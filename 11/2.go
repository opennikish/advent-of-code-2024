package main

import (
	"bytes"
	"fmt"
	"iter"
	"os"
	"strconv"
)

func main() {
	res := solve("in.txt", 75)
	fmt.Println(res)
}

func solve(path string, blinks int) int {
	bs, err := os.ReadFile(path)
	check(err)
	bs = bytes.TrimSpace(bs)
	stones := bytesToInts(bs)

	return transform(stones, blinks)
}

func bytesToInts(bs []byte) []int {
	if bs == nil {
		return nil
	}
	parts := bytes.Split(bs, []byte{' '})
	res := make([]int, len(parts))
	for i, p := range parts {
		x, err := strconv.Atoi(string(p))
		check(err)
		res[i] = x
	}
	return res
}

func transform(stones []int, blinks int) int {
	curr := map[int]int{}
	for _, s := range stones {
		curr[s]++
	}

	next := map[int]int{}
	for blinks > 0 {
		for stone, count := range curr {
			if stone == 0 {
				next[1] += count
				continue
			}

			len := intLen(stone)
			if len%2 == 0 {
				left, right := split(stone, len)
				next[left] += count
				next[right] += count
				continue
			}

			next[stone*2024] += count
		}

		curr = next
		next = map[int]int{}
		blinks--
	}

	sum := 0
	for _, count := range curr {
		sum += count
	}

	return sum
}

func iterToSlice(it iter.Seq[int]) []int {
	xs := []int{}
	for x := range it {
		xs = append(xs, x)
	}
	return xs
}

// split splits number into two numbers, one containg left half on digits and second right half
// examples: 123123 => 123 123; 123001 => 123, 1
func split(stone, len int) (int, int) {
	if len%2 != 0 {
		panic("not even stone length")
	}

	left, right := stone, 0
	exp := 1

	for i := 0; i < len/2; i++ {
		rem := left % 10
		right += rem * exp

		left /= 10
		exp *= 10
	}

	return left, right
}

func intLen(x int) int {
	l := 0
	for x > 0 {
		x /= 10
		l++
	}
	return l
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
