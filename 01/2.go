package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var noWhiteSpaceReg = regexp.MustCompile(`\S+`)

func main() {
	nums1 := []int{}
	nums2 := map[int]int{}

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
		nums2[n2]++
	})

	sum := 0
	for _, n := range nums1 {
		sum += n * nums2[n]
	}

	fmt.Println(sum)
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
