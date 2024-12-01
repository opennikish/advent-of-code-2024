package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var noWhiteSpaceReg = regexp.MustCompile(`\S+`)

func main() {
	nums1, nums2 := []int{}, []int{}

	readLines("input1.txt", func(line string) {
		parts := noWhiteSpaceReg.FindAllString(line, -1)

		n1, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		n2, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		nums1 = append(nums1, n1)
		nums2 = append(nums2, n2)
	})

	slices.Sort(nums1)
	slices.Sort(nums2)

	sum := 0
	for i, n := range nums1 {
		sum += abs(n - nums2[i])
	}

	fmt.Println(sum)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func readLines(path string, fn func(line string)) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		fn(line)
	}
}
