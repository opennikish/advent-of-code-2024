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

	antenaMap2 := make([][]byte, len(antenaMap))
	for i := range antenaMap {
		antenaMap2[i] = append([]byte(nil), antenaMap[i]...)
	}

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
	x, y := abs(p1.x-p2.x), p1.y-p2.y
	n, m := len(antenaMap), len(antenaMap[0])

	count := 0
	if antenaMap[p1.x][p1.y] != '#' {
		count++
	}
	antenaMap[p1.x][p1.y] = '#'

	if antenaMap[p2.x][p2.y] != '#' {
		count++
	}
	antenaMap[p2.x][p2.y] = '#'

	// go up
	for a1x, a1y := p1.x-x, p1.y+y; a1x >= 0 && a1x < n && a1y >= 0 && a1y < m; {
		if antenaMap[a1x][a1y] != '#' {
			count++
			antenaMap[a1x][a1y] = '#'
		}
		a1x -= x
		a1y += y
	}

	// go down
	for a2x, a2y := p2.x+x, p2.y-y; a2x >= 0 && a2x < n && a2y >= 0 && a2y < m; {
		if antenaMap[a2x][a2y] != '#' {
			count++
			antenaMap[a2x][a2y] = '#'
		}
		a2x += x
		a2y -= y
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
