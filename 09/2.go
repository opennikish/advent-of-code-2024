package main

import (
	"bytes"
	"fmt"
	"os"
)

type slot struct {
	i, size, id int
}

func main() {
	res := solve("in.txt")
	fmt.Println(res)
}

func solve(path string) int {
	bs, err := os.ReadFile(path)
	check(err)
	fmap := bytes.TrimSpace(bs)

	fileSlots, emptySlots := collect(fmap)

	for i := len(fileSlots) - 1; i >= 0; i-- {
		file := fileSlots[i]

		for j, empty := range emptySlots {
			if empty.i > file.i {
				emptySlots = emptySlots[0:j]
				break
			}

			if empty.size < file.size {
				continue
			}

			pos := empty.i
			empty.i += file.size
			empty.size -= file.size
			file.i = pos

			break
		}

		if len(emptySlots) == 0 {
			break
		}
	}

	return checksum(fileSlots)
}

func checksum(fileSlots []*slot) int {
	s := 0
	for _, file := range fileSlots {
		for k := file.i; k < file.i+file.size; k++ {
			s += file.id * k
		}
	}
	return s
}

func collect(fmap []byte) ([]*slot, []*slot) {
	fileSlots, emptySlots := []*slot{}, []*slot{}
	id, pos := 0, 0
	for i, b := range fmap {
		size := int(b - '0')
		if (i+1)%2 == 0 {
			emptySlots = append(emptySlots, &slot{i: pos, size: size})
		} else {
			fileSlots = append(fileSlots, &slot{i: pos, size: size, id: id})
			id++
		}
		pos += size
	}

	return fileSlots, emptySlots
}

func print(fileSlots []*slot) {
	for _, s := range fileSlots {
		fmt.Printf("(%d, %d, %d) ", s.i, s.size, s.id)
	}
	fmt.Println()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
