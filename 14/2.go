package main

import (
	"adventofcode2024/lib"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type Command int

const (
	Pause Command = iota
	Next
	Prev
	Quit
)

var cmdNames = map[Command]string{
	Pause: "pause",
	Next:  "next",
	Prev:  "prev",
	Quit:  "quit",
}

func (c Command) String() string {
	return cmdNames[c]
}

var rNums = regexp.MustCompile(`[\-0-9]+`)

type robot struct {
	x, y       int
	movX, movY int
}

type point struct {
	x, y int
}

const width = 101
const height = 103

func main() {
	input, err := lib.GetInput(14)
	lib.Check(err)
	robots := parseRobots(input)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cmds, errc := cmdReader(ctx)

	ticks := time.Tick(1000 * time.Millisecond)

	paused := false
	count := 0
	cache := map[point]struct{}{}

	grid := make([][]byte, height)
	for i := range grid {
		grid[i] = make([]byte, width)
	}

	for {
		select {
		case cmd := <-cmds:
			log("cmd: %s", cmd)

			if cmd == Quit {
				fmt.Println("Bye")
				return
			}

			if cmd == Pause {
				paused = paused != true
			}

		case <-ticks:
			log("tick")
			if paused {
				continue
			}
			clearScreen()
			count++
			move(robots, width, height)
			cacheRobots(cache, robots)
			render(cache, grid, count)

		case <-ctx.Done():
			fmt.Println("Bye")
			return
		case err := <-errc:
			fmt.Println(err)
			return
		}
	}
}

func log(format string, a ...any) {
	if len(a) == 0 {
		fmt.Fprint(os.Stderr, format+"\n")
	} else {
		fmt.Fprintf(os.Stderr, format+"\n", a)
	}
}

func move(robots []*robot, width, height int) {
	transform := func(x, move, limit int) int {
		steps := abs(move) % limit
		if move > 0 {
			return (x + steps) % limit
		}

		x -= steps
		if x < 0 {
			x += limit
		}
		return x
	}

	for _, r := range robots {
		r.x = transform(r.x, r.movX, width)
		r.y = transform(r.y, r.movY, height)
	}
}

func cacheRobots(cache map[point]struct{}, robots []*robot) {
	clear(cache)
	for _, r := range robots {
		cache[point{x: r.x, y: r.y}] = struct{}{}
	}
}

func render(robots map[point]struct{}, grid [][]byte, count int) {
	fmt.Printf("--- SEC: %d ---\n\n", count)

	for i := range height {
		for j := range width {
			if _, ok := robots[point{y: i, x: j}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func cmdReader(ctx context.Context) (<-chan Command, <-chan error) {
	// disable input buffering
	err := exec.Command("stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
	lib.Check(err)

	// do not display entered characters on the screen
	err = exec.Command("stty", "-f", "/dev/tty", "-echo").Run()
	lib.Check(err)

	cmds, errc := make(chan Command), make(chan error, 1)
	buf := make([]byte, 3)

	cmdMap := map[string]Command{
		" ":                        Pause,
		string([]byte{27, 91, 67}): Next, // escape sequence for right arrow
		string([]byte{27, 91, 68}): Prev, // escape sequence for left arrow
		"q":                        Quit,
	}

	go func() {
		for {
			if ctx.Err() != nil {
				close(errc)
				close(cmds)
				return
			}

			n, err := os.Stdin.Read(buf)
			if err != nil {
				errc <- fmt.Errorf("read stdin: %w", err)
				close(errc)
				close(cmds)
			}

			if cmd, ok := cmdMap[string(buf[:n])]; ok {
				cmds <- cmd
			}
		}
	}()

	return cmds, errc
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func parseRobots(input []byte) []*robot {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	robots := make([]*robot, len(lines))
	for i, l := range lines {
		nums := rNums.FindAllString(l, -1)
		if len(nums) != 4 {
			panic("corrupted robot: " + l)
		}
		robots[i] = &robot{
			x:    toInt(nums[0]),
			y:    toInt(nums[1]),
			movX: toInt(nums[2]),
			movY: toInt(nums[3]),
		}
	}
	return robots
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	lib.Check(err)
	return n
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
