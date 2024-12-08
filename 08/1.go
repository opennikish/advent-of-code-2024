package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	res := solve("in.txt")
	fmt.Println(res)
}

type point struct {
	x, y int
}

func solve(path string) int {
	bs, err := os.ReadFile(path)
	check(err)
	bs = bytes.TrimSpace(bs)
	antenaMap := bytes.Split(bs, []byte{'\n'})

	coordsByFreq := map[byte][]point{}
	for i := range antenaMap {
		for j := range antenaMap[0] {
			if antenaMap[i][j] != '.' {
				freq := antenaMap[i][j]
				coordsByFreq[freq] = append(coordsByFreq[freq], point{x: i, y: j})
			}
		}
	}

	total := 0
	for _, coords := range coordsByFreq {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				total += calcAntiNodes(antenaMap, coords[i], coords[j])
			}
		}
	}

	return total
}

// p1 is always higher (in x coord) than p2, so we could simplify calculations
func calcAntiNodes(antenaMap [][]byte, p1, p2 point) int {
	x, y := abs(p1.x-p2.x), abs(p1.y-p2.y)

	var a1x, a1y int
	a1x = p1.x - x
	if p1.y > p2.y {
		a1y = p1.y + y // slope: /
	} else {
		a1y = p1.y - y // slope: \
	}

	var a2x, a2y int
	a2x = p2.x + x
	if p2.y < p1.y {
		a2y = p2.y - y // slope: /
	} else {
		a2y = p2.y + y // slope: \
	}

	n, m := len(antenaMap), len(antenaMap[0])
	count := 0

	if a1x >= 0 && a1x < n && a1y >= 0 && a1y < m && antenaMap[a1x][a1y] != '#' {
		antenaMap[a1x][a1y] = '#'
		count++
	}

	if a2x >= 0 && a2x < n && a2y >= 0 && a2y < m && antenaMap[a2x][a2y] != '#' {
		antenaMap[a2x][a2y] = '#'
		count++
	}

	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
